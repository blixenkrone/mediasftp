

# MEDIA SFTP

### This contains the use case, requirements and specification of the FTP server for uploading content

It should:
- implement the SSH FTP (SFTP)
- be accessed through a DNS host and port i.e: ftp.byrd.news:21.
- contain login information
- have web interface for adding/removing/downloading stuff
- have user delegation (admin/profiles) if loginID==allowedUserID 
- Be deployed and reached over HTTP TCP (digital ocean)
- Storing the content in digital ocean
- use openSSH on linux?
- Support bursts/spikes in workload

### FTP lib: https://github.com/jlaffaye/ftp


