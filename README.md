# Fudge

I had some success with on Phishing engagements by base64 encoding the contents of a file, putting it into a HTML file and having it decode and drop when the user opened the HTML file. 


# Usage
There are precompiled binaries available on the [releases](https://github.com/dale-ruane/Fudge/releases) tab for both **Windows 64-bit** and **Linux 64-bit**. As they're Go binaries you shouldn't need anything except to mark them executable and run them.

If you have Go set up on your system you can download the latest version from the repo with:

    go get github.com/dale-ruane/Fudge

## Command line reference
      -n string
            Name of downloaded file (default "nastypasty.exe") (This is the filename the target will see when the file pops out of the HTML file)
      -s string
            Source File (default "test.exe")

Use the above flags to set the file to encode, and the filename which will be given to the user when they open the HTML file.  

## Example usage

    ./Fudge -s implant.exe -n passwordchecker.exe
The above will produce  **output.html** which will contain implant.exe and when the user opens the HTML file it will be named passwordchecker.exe. 

# Supported Files
As far as I am aware, any source file should work. As Fudge uses the raw bytes of the original file, they should be identical when they land on the target host.

The main consideration being that the input and output file extensions should be the same if you wish the file to work the same on the target host.

If you try any files and they dont work, please open a [issue](https://github.com/dale-ruane/Fudge/issues) and provide an example file if possible.

# Disclaimer
This tool is released purely for educational purposes, or for use on penetration testing engagement where appropriate authorisation has been obtained. Any usage outside of this is entirely at the user's own risk.
