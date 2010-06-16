#
# Ruby scripts that parese the XML prototypes of the XCB library and
# generates go language from it.
# Yes, I don't like Python too much. :p
#
require "rexml/document"

TEST_FILE = 'xcb/xproto.xml'
PAD       = 28

# Gets an attribute from the XML element 
def get_attr(el, attr)
  return el.attributes[attr]
end

 

# Returns true if a name is all uppercase
def all_upcase?(name)
  return name == name.upcase
end

# Gets the name. Corrects it so itt's in line with Go style, mostly.  
def get_name(el)
  name = get_attr(el, "name")
  return name if !name # no name, so nothing else to do
  return name if all_upcase?(name) # all upcase name is fine
  newname = name.split('_').map() do |val|
    val[0] = val[0].upcase
    val
  end.join
  # Split on any underscores, capitalize, and join again.   
  return newname
end

# Gets the type of a list of such. Gandles problematical cases like 
# void
def get_type(el)
  type = get_attr(el, 'type')
  return type if !type # no type, so nothing else to do
  return '[]BYTE' if type == 'void' 
  return type
end

# outputs a type definition without a struct
def type(out, old, new) 
  out.puts "type #{old.ljust(PAD)}   #{new}"
end

# outputs a type definition with a struct  
def struct(out, name, fields) 
  out.puts "\ntype #{name.ljust(PAD-2)} struct {"
  for field in fields do
    out.puts field 
  end
  out.puts "}\n\n"
end

# generates a field line for go from field element.
def make_field(el)
  name = get_name(el)
  type = get_type(el)
  comm = ""
  enum =  get_attr(el, "enum")
  comm = "  // Use constants starting with #{enum}" if enum 
  return "  #{name.ljust(PAD)}  #{type}#{comm}"
end
 
# generates a padding line for go from padding element.
def make_padding(el, count)
  size = get_attr(el, "bytes")
  name = "Padding#{count}"  
  return "  #{name.to_s.ljust(PAD-3)}  [#{size}]PADDING"
end

# generates a field line for go from a list element.
def make_list(el)
  name    = get_name(el)
  type    = get_type(el)
  value   = get_value(el)
  # value   = get_attr(el, 'bytes')
     
  if type=='char' # list of char -> STRING8
    return "  #{name.ljust(PAD)}  STRING8"
  end
  return "  #{name.ljust(PAD-2)}  [#{value}]#{type}"
end
 
# Processes structs
def process_struct(el, out, name = nil, extra_fields = nil)
  name  ||= get_name(el)
  pad     = 1 # amount of padding field added
  fields  = [] # field data
  if extra_fields # add extra fields first
    fields += extra_fields 
  end
  el.elements.each do | field |
    subname = field.name
    case subname
      when 'field' 
        fields << make_field(field)
      when 'pad'
        fields << make_padding(field, pad)
        pad    += 1
      when 'list'
        fields << make_list(field)
      when 'reply' # process replies for request structs
        process_reply(el, field, out)
      else 
        warn "Unknown sub-element #{subname} #{field} in a struct"
    end
  end  
  struct(out, name, fields)
      
end

def process_xidtype(el, out)
  name = get_name(el)
  type(out, name, 'XID')
end

def process_xidunion(el, out)
  name = get_name(el)
  type(out, name, 'XID // Was a union of XID types.')
end

def process_typedef(el, out)
  oldname = get_attr(el,"oldname")
  newname = get_attr(el,"newname")
  type(out, newname, oldname)
end


def make_item_value(sub) 
    subname = sub.name
    case subname 
      when 'value'
        return sub.text.strip.to_i.to_s
      when 'bit'
        val = 1 << sub.text.strip.to_i
        hex = "0x" + val.to_s(16).rjust(8,'0')
        return hex
      else
        warn "Unknown value: #{sub}"
        return nil
    end  
end

# Retuns the value or bit value for this element  
def get_value(el)
  el.elements.each do  |sub|
    val = make_item_value(sub)
    return val if val
  end
  return nil  
end

# 
def make_item(el, enumname)
  name = get_name(el)
  val  = get_value(el)
  unless val 
    # if the value is missing, generate it and hope it's OK
    unless @last_value
      warning = "// Value and last value missing here for #{name}!"
      warn warning
      return warning
    end
    if @last_value.slice(0..1) == "0x"
      val   = @last_value.to_i(16)
    else   
      val   = @last_value.to_i
    end
    val  += 1 # one more than last value
    val   = val.to_s
  end
  @last_value = val.to_s
  # Make the enum name similar to the individual name style
  wholename   = enumname + name
  if all_upcase?(name)
    wholename = enumname.upcase + '_' + name
  end
  
  return "  #{wholename.ljust(PAD)}  = #{val}"    
end
 

# outputs an enum definition
def enum(out, name, items)
  # out.puts "\ntype #{name.ljust(PAD)} #{type}"
  out.puts "\nconst ("
  for item in items do
    out.puts item
  end
  out.puts ")\n"
end

def process_enum(el, out)
  # reset last_value
  @last_value = nil
  name    = get_name(el)
  pad     = 1 # amount of padding field added
  items   = [] # item data  
  el.elements.each do | item |
    subname = item.name
    case subname
      when 'item'
        items << make_item(item, name)
      else 
        warn "Unknown sub-element #{subname} #{field} in an item"
    end
  end  
  enum(out, name, items)
  # reset last_value again
  @last_value = nil
end

# Adds a single constant declaration to the output
def const(out, name, value)
  out.puts "const #{name.ljust(PAD-1)}  = #{value}"
end

def process_union(el, out)
  last = nil
  el.elements.each do | item |
    last = make_field(item)
  end
  out.puts "// This struct is generated from a union. "
  out.puts "// Only the last field is added. Probably it will not work."
  if last    
    struct(out, get_name(el), [last])      
  end
end

def process_event(el, out)
  name      = get_name(el) + 'Event'
  process_struct(el, out, name)
  constname = get_name(el) + 'EventCode'
  constval  = get_attr(el, 'number')
  const(out, constname, constval)
  out.puts  
  @events ||= {} # must keep track of events for eventcopy 
  @events[get_name(el)] = el 
end

def process_eventcopy(el, out)
  name      = get_name(el) + 'Event'
  ref       = get_attr(el, 'ref')
  copied    = @events[ref]
  if !copied
    warn "Cannot coppy event from #{ref} to #{name}"
    return
  end
  process_struct(copied, out, name)
  constname = get_name(el) + 'EventCode'
  constval  = get_attr(el, 'number')
  const(out, constname, constval)  
end

# Processes replies for a request.
def process_reply(request, reply, out) 
  name     = get_name(request) + 'Reply'
  process_struct(reply, out, name) 
  # Recurse once more
end

def process_request(el, out)
  name      = get_name(el) + 'Request'
  extra     = "  Opcode".ljust(PAD) + "    CARD8"
  process_struct(el, out, name, [extra])
  constname = get_name(el) + 'Opcode'
  constval  = get_attr(el, 'opcode')
  const(out, constname, constval)
  out.puts  
end

def process_error(el, out)
  name      = get_name(el) + 'Error'
  process_struct(el, out, name) 
  constname = get_name(el) + 'ErrorCode'
  constval  = get_attr(el, 'number')
  const(out, constname, constval)
  out.puts
  @errors ||= {} # must keep track of errors for errorcopy 
  @errors[get_name(el)] = el 
end

def process_errorcopy(el, out)
  name      = get_name(el) + 'Error'
  ref       = get_attr(el, 'ref')
  copied    = @errors[ref]
  if !copied
    warn "Cannot coppy error from #{ref} to #{name}"
    return
  end
  process_struct(copied, out, name)
  constname = get_name(el) + 'ErrorCode'
  constval  = get_attr(el, 'number')
  const(out, constname, constval)
  
end


def define_basics(out)
  basics = <<-END_OF_BASICS
  package x
  
  // Generated Types Follow below this line 
  
END_OF_BASICS
  basics = basics.split("\n").map(&:strip).join("\n")
  out.puts(basics)
end


def process_root(xroot, out) 
  define_basics(out)
  xroot.elements.each do | element |
    name = element.name.downcase.strip
    el   = element
    case name 
      when "struct"
        process_struct(el, out)  
      when "xidtype"  
        process_xidtype(el, out)
      when "xidunion" 
        process_xidunion(el, out)
      when "typedef"  
        process_typedef(el, out)
      when "enum"     
        process_enum(el, out)
      when "union"    
        process_union(el, out)
      when "event"    
        process_event(el, out)
      when "eventcopy"
        process_eventcopy(el, out)
      when "request"
        process_request(el, out)
      when "error"    
        process_error(el, out)
      when "errorcopy"
        process_errorcopy(el, out)
      else 
        warning = "/* FIXME: Unknown element:\n #{element}\n */"
        warn warning
        out << warning
    end
  end  
end


def main(args)
  @last_value = nil
  infilename  = args[0] || TEST_FILE 
  data        = File.read(infilename)
  xml         = REXML::Document.new(data)
  xroot       = xml.root
  outname     = xroot.attributes["header"]
  outfilename = "#{outname}.go"
  outfile    = File.open(outfilename,"w+")
  p outfilename
   
  process_root(xroot, outfile)
  outfile.close
end


main(ARGV)










