## GeoSnitch

A proof of concept [osquery](https://github.com/osquery/osquery) extension to determine a user's current physical location by municipality, based-on the device's wifi site survey and calculated by Google.  

This system was designed to determine if a user accessing a FedRamp environment was respecting the geography-based firewall rule by employing a VPN.  

Tested successfully on Windows 10 and Pop!OS.  
MacOS is _not supported_ due to Apple sterilizing location information.  

### Running this thing

Start osquery as `osqueryi --extensions_socket=~/.osquery/shell.em`  
_THEN_ start daemon as `./GeoSnitch --socket ~/.osquery/shell.em`  
