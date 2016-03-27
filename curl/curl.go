package curl
// #include <curl/curl.h>
import "C"

// CURLOPT_XXX
const (
    CURLOPT_WRITEDATA                = C.CURLOPT_WRITEDATA
    CURLOPT_URL                      = C.CURLOPT_URL
    CURLOPT_PORT                     = C.CURLOPT_PORT
    CURLOPT_PROXY                    = C.CURLOPT_PROXY
    CURLOPT_USERPWD                  = C.CURLOPT_USERPWD
    CURLOPT_PROXYUSERPWD             = C.CURLOPT_PROXYUSERPWD
    CURLOPT_RANGE                    = C.CURLOPT_RANGE
    CURLOPT_READDATA                 = C.CURLOPT_READDATA
    CURLOPT_ERRORBUFFER              = C.CURLOPT_ERRORBUFFER
    CURLOPT_WRITEFUNCTION            = C.CURLOPT_WRITEFUNCTION
    CURLOPT_READFUNCTION             = C.CURLOPT_READFUNCTION
    CURLOPT_TIMEOUT                  = C.CURLOPT_TIMEOUT
    CURLOPT_INFILESIZE               = C.CURLOPT_INFILESIZE
    CURLOPT_POSTFIELDS               = C.CURLOPT_POSTFIELDS
    CURLOPT_REFERER                  = C.CURLOPT_REFERER
    CURLOPT_FTPPORT                  = C.CURLOPT_FTPPORT
    CURLOPT_USERAGENT                = C.CURLOPT_USERAGENT
    CURLOPT_LOW_SPEED_LIMIT          = C.CURLOPT_LOW_SPEED_LIMIT
    CURLOPT_LOW_SPEED_TIME           = C.CURLOPT_LOW_SPEED_TIME
    CURLOPT_RESUME_FROM              = C.CURLOPT_RESUME_FROM
    CURLOPT_COOKIE                   = C.CURLOPT_COOKIE
    CURLOPT_HTTPHEADER               = C.CURLOPT_HTTPHEADER
    CURLOPT_HTTPPOST                 = C.CURLOPT_HTTPPOST
    CURLOPT_SSLCERT                  = C.CURLOPT_SSLCERT
    CURLOPT_KEYPASSWD                = C.CURLOPT_KEYPASSWD
    CURLOPT_CRLF                     = C.CURLOPT_CRLF
    CURLOPT_QUOTE                    = C.CURLOPT_QUOTE
    CURLOPT_HEADERDATA               = C.CURLOPT_HEADERDATA
    CURLOPT_COOKIEFILE               = C.CURLOPT_COOKIEFILE
    CURLOPT_SSLVERSION               = C.CURLOPT_SSLVERSION
    CURLOPT_TIMECONDITION            = C.CURLOPT_TIMECONDITION
    CURLOPT_TIMEVALUE                = C.CURLOPT_TIMEVALUE
    CURLOPT_CUSTOMREQUEST            = C.CURLOPT_CUSTOMREQUEST
    CURLOPT_STDERR                   = C.CURLOPT_STDERR
    CURLOPT_POSTQUOTE                = C.CURLOPT_POSTQUOTE
    CURLOPT_OBSOLETE40               = C.CURLOPT_OBSOLETE40
    CURLOPT_VERBOSE                  = C.CURLOPT_VERBOSE
    CURLOPT_HEADER                   = C.CURLOPT_HEADER
    CURLOPT_NOPROGRESS               = C.CURLOPT_NOPROGRESS
    CURLOPT_NOBODY                   = C.CURLOPT_NOBODY
    CURLOPT_FAILONERROR              = C.CURLOPT_FAILONERROR
    CURLOPT_UPLOAD                   = C.CURLOPT_UPLOAD
    CURLOPT_POST                     = C.CURLOPT_POST
    CURLOPT_DIRLISTONLY              = C.CURLOPT_DIRLISTONLY
    CURLOPT_APPEND                   = C.CURLOPT_APPEND
    CURLOPT_NETRC                    = C.CURLOPT_NETRC
    CURLOPT_FOLLOWLOCATION           = C.CURLOPT_FOLLOWLOCATION
    CURLOPT_TRANSFERTEXT             = C.CURLOPT_TRANSFERTEXT
    CURLOPT_PUT                      = C.CURLOPT_PUT
    CURLOPT_PROGRESSFUNCTION         = C.CURLOPT_PROGRESSFUNCTION
    CURLOPT_PROGRESSDATA             = C.CURLOPT_PROGRESSDATA
    CURLOPT_AUTOREFERER              = C.CURLOPT_AUTOREFERER
    CURLOPT_PROXYPORT                = C.CURLOPT_PROXYPORT
    CURLOPT_POSTFIELDSIZE            = C.CURLOPT_POSTFIELDSIZE
    CURLOPT_HTTPPROXYTUNNEL          = C.CURLOPT_HTTPPROXYTUNNEL
    CURLOPT_INTERFACE                = C.CURLOPT_INTERFACE
    CURLOPT_KRBLEVEL                 = C.CURLOPT_KRBLEVEL
    CURLOPT_SSL_VERIFYPEER           = C.CURLOPT_SSL_VERIFYPEER
    CURLOPT_CAINFO                   = C.CURLOPT_CAINFO
    CURLOPT_MAXREDIRS                = C.CURLOPT_MAXREDIRS
    CURLOPT_FILETIME                 = C.CURLOPT_FILETIME
    CURLOPT_TELNETOPTIONS            = C.CURLOPT_TELNETOPTIONS
    CURLOPT_MAXCONNECTS              = C.CURLOPT_MAXCONNECTS
    CURLOPT_OBSOLETE72               = C.CURLOPT_OBSOLETE72
    CURLOPT_FRESH_CONNECT            = C.CURLOPT_FRESH_CONNECT
    CURLOPT_FORBID_REUSE             = C.CURLOPT_FORBID_REUSE
    CURLOPT_RANDOM_FILE              = C.CURLOPT_RANDOM_FILE
    CURLOPT_EGDSOCKET                = C.CURLOPT_EGDSOCKET
    CURLOPT_CONNECTTIMEOUT           = C.CURLOPT_CONNECTTIMEOUT
    CURLOPT_HEADERFUNCTION           = C.CURLOPT_HEADERFUNCTION
    CURLOPT_HTTPGET                  = C.CURLOPT_HTTPGET
    CURLOPT_SSL_VERIFYHOST           = C.CURLOPT_SSL_VERIFYHOST
    CURLOPT_COOKIEJAR                = C.CURLOPT_COOKIEJAR
    CURLOPT_SSL_CIPHER_LIST          = C.CURLOPT_SSL_CIPHER_LIST
    CURLOPT_HTTP_VERSION             = C.CURLOPT_HTTP_VERSION
    CURLOPT_FTP_USE_EPSV             = C.CURLOPT_FTP_USE_EPSV
    CURLOPT_SSLCERTTYPE              = C.CURLOPT_SSLCERTTYPE
    CURLOPT_SSLKEY                   = C.CURLOPT_SSLKEY
    CURLOPT_SSLKEYTYPE               = C.CURLOPT_SSLKEYTYPE
    CURLOPT_SSLENGINE                = C.CURLOPT_SSLENGINE
    CURLOPT_SSLENGINE_DEFAULT        = C.CURLOPT_SSLENGINE_DEFAULT
    CURLOPT_DNS_USE_GLOBAL_CACHE     = C.CURLOPT_DNS_USE_GLOBAL_CACHE
    CURLOPT_DNS_CACHE_TIMEOUT        = C.CURLOPT_DNS_CACHE_TIMEOUT
    CURLOPT_PREQUOTE                 = C.CURLOPT_PREQUOTE
    CURLOPT_DEBUGFUNCTION            = C.CURLOPT_DEBUGFUNCTION
    CURLOPT_DEBUGDATA                = C.CURLOPT_DEBUGDATA
    CURLOPT_COOKIESESSION            = C.CURLOPT_COOKIESESSION
    CURLOPT_CAPATH                   = C.CURLOPT_CAPATH
    CURLOPT_BUFFERSIZE               = C.CURLOPT_BUFFERSIZE
    CURLOPT_NOSIGNAL                 = C.CURLOPT_NOSIGNAL
    CURLOPT_SHARE                    = C.CURLOPT_SHARE
    CURLOPT_PROXYTYPE                = C.CURLOPT_PROXYTYPE
    CURLOPT_ACCEPT_ENCODING          = C.CURLOPT_ACCEPT_ENCODING
    CURLOPT_PRIVATE                  = C.CURLOPT_PRIVATE
    CURLOPT_HTTP200ALIASES           = C.CURLOPT_HTTP200ALIASES
    CURLOPT_UNRESTRICTED_AUTH        = C.CURLOPT_UNRESTRICTED_AUTH
    CURLOPT_FTP_USE_EPRT             = C.CURLOPT_FTP_USE_EPRT
    CURLOPT_HTTPAUTH                 = C.CURLOPT_HTTPAUTH
    CURLOPT_SSL_CTX_FUNCTION         = C.CURLOPT_SSL_CTX_FUNCTION
    CURLOPT_SSL_CTX_DATA             = C.CURLOPT_SSL_CTX_DATA
    CURLOPT_FTP_CREATE_MISSING_DIRS  = C.CURLOPT_FTP_CREATE_MISSING_DIRS
    CURLOPT_PROXYAUTH                = C.CURLOPT_PROXYAUTH
    CURLOPT_FTP_RESPONSE_TIMEOUT     = C.CURLOPT_FTP_RESPONSE_TIMEOUT
    CURLOPT_IPRESOLVE                = C.CURLOPT_IPRESOLVE
    CURLOPT_MAXFILESIZE              = C.CURLOPT_MAXFILESIZE
    CURLOPT_INFILESIZE_LARGE         = C.CURLOPT_INFILESIZE_LARGE
    CURLOPT_RESUME_FROM_LARGE        = C.CURLOPT_RESUME_FROM_LARGE
    CURLOPT_MAXFILESIZE_LARGE        = C.CURLOPT_MAXFILESIZE_LARGE
    CURLOPT_NETRC_FILE               = C.CURLOPT_NETRC_FILE
    CURLOPT_USE_SSL                  = C.CURLOPT_USE_SSL
    CURLOPT_POSTFIELDSIZE_LARGE      = C.CURLOPT_POSTFIELDSIZE_LARGE
    CURLOPT_TCP_NODELAY              = C.CURLOPT_TCP_NODELAY
    CURLOPT_FTPSSLAUTH               = C.CURLOPT_FTPSSLAUTH
    CURLOPT_IOCTLFUNCTION            = C.CURLOPT_IOCTLFUNCTION
    CURLOPT_IOCTLDATA                = C.CURLOPT_IOCTLDATA
    CURLOPT_FTP_ACCOUNT              = C.CURLOPT_FTP_ACCOUNT
    CURLOPT_COOKIELIST               = C.CURLOPT_COOKIELIST
    CURLOPT_IGNORE_CONTENT_LENGTH    = C.CURLOPT_IGNORE_CONTENT_LENGTH
    CURLOPT_FTP_SKIP_PASV_IP         = C.CURLOPT_FTP_SKIP_PASV_IP
    CURLOPT_FTP_FILEMETHOD           = C.CURLOPT_FTP_FILEMETHOD
    CURLOPT_LOCALPORT                = C.CURLOPT_LOCALPORT
    CURLOPT_LOCALPORTRANGE           = C.CURLOPT_LOCALPORTRANGE
    CURLOPT_CONNECT_ONLY             = C.CURLOPT_CONNECT_ONLY
    CURLOPT_CONV_FROM_NETWORK_FUNCTION = C.CURLOPT_CONV_FROM_NETWORK_FUNCTION
    CURLOPT_CONV_TO_NETWORK_FUNCTION = C.CURLOPT_CONV_TO_NETWORK_FUNCTION
    CURLOPT_CONV_FROM_UTF8_FUNCTION  = C.CURLOPT_CONV_FROM_UTF8_FUNCTION
    CURLOPT_MAX_SEND_SPEED_LARGE     = C.CURLOPT_MAX_SEND_SPEED_LARGE
    CURLOPT_MAX_RECV_SPEED_LARGE     = C.CURLOPT_MAX_RECV_SPEED_LARGE
    CURLOPT_FTP_ALTERNATIVE_TO_USER  = C.CURLOPT_FTP_ALTERNATIVE_TO_USER
    CURLOPT_SOCKOPTFUNCTION          = C.CURLOPT_SOCKOPTFUNCTION
    CURLOPT_SOCKOPTDATA              = C.CURLOPT_SOCKOPTDATA
    CURLOPT_SSL_SESSIONID_CACHE      = C.CURLOPT_SSL_SESSIONID_CACHE
    CURLOPT_SSH_AUTH_TYPES           = C.CURLOPT_SSH_AUTH_TYPES
    CURLOPT_SSH_PUBLIC_KEYFILE       = C.CURLOPT_SSH_PUBLIC_KEYFILE
    CURLOPT_SSH_PRIVATE_KEYFILE      = C.CURLOPT_SSH_PRIVATE_KEYFILE
    CURLOPT_FTP_SSL_CCC              = C.CURLOPT_FTP_SSL_CCC
    CURLOPT_TIMEOUT_MS               = C.CURLOPT_TIMEOUT_MS
    CURLOPT_CONNECTTIMEOUT_MS        = C.CURLOPT_CONNECTTIMEOUT_MS
    CURLOPT_HTTP_TRANSFER_DECODING   = C.CURLOPT_HTTP_TRANSFER_DECODING
    CURLOPT_HTTP_CONTENT_DECODING    = C.CURLOPT_HTTP_CONTENT_DECODING
    CURLOPT_NEW_FILE_PERMS           = C.CURLOPT_NEW_FILE_PERMS
    CURLOPT_NEW_DIRECTORY_PERMS      = C.CURLOPT_NEW_DIRECTORY_PERMS
    CURLOPT_POSTREDIR                = C.CURLOPT_POSTREDIR
    CURLOPT_SSH_HOST_PUBLIC_KEY_MD5  = C.CURLOPT_SSH_HOST_PUBLIC_KEY_MD5
    CURLOPT_OPENSOCKETFUNCTION       = C.CURLOPT_OPENSOCKETFUNCTION
    CURLOPT_OPENSOCKETDATA           = C.CURLOPT_OPENSOCKETDATA
    CURLOPT_COPYPOSTFIELDS           = C.CURLOPT_COPYPOSTFIELDS
    CURLOPT_PROXY_TRANSFER_MODE      = C.CURLOPT_PROXY_TRANSFER_MODE
    CURLOPT_SEEKFUNCTION             = C.CURLOPT_SEEKFUNCTION
    CURLOPT_SEEKDATA                 = C.CURLOPT_SEEKDATA
    CURLOPT_CRLFILE                  = C.CURLOPT_CRLFILE
    CURLOPT_ISSUERCERT               = C.CURLOPT_ISSUERCERT
    CURLOPT_ADDRESS_SCOPE            = C.CURLOPT_ADDRESS_SCOPE
    CURLOPT_CERTINFO                 = C.CURLOPT_CERTINFO
    CURLOPT_USERNAME                 = C.CURLOPT_USERNAME
    CURLOPT_PASSWORD                 = C.CURLOPT_PASSWORD
    CURLOPT_PROXYUSERNAME            = C.CURLOPT_PROXYUSERNAME
    CURLOPT_PROXYPASSWORD            = C.CURLOPT_PROXYPASSWORD
    CURLOPT_NOPROXY                  = C.CURLOPT_NOPROXY
    CURLOPT_TFTP_BLKSIZE             = C.CURLOPT_TFTP_BLKSIZE
    CURLOPT_SOCKS5_GSSAPI_SERVICE    = C.CURLOPT_SOCKS5_GSSAPI_SERVICE
    CURLOPT_SOCKS5_GSSAPI_NEC        = C.CURLOPT_SOCKS5_GSSAPI_NEC
    CURLOPT_PROTOCOLS                = C.CURLOPT_PROTOCOLS
    CURLOPT_REDIR_PROTOCOLS          = C.CURLOPT_REDIR_PROTOCOLS
    CURLOPT_SSH_KNOWNHOSTS           = C.CURLOPT_SSH_KNOWNHOSTS
    CURLOPT_SSH_KEYFUNCTION          = C.CURLOPT_SSH_KEYFUNCTION
    CURLOPT_SSH_KEYDATA              = C.CURLOPT_SSH_KEYDATA
    CURLOPT_MAIL_FROM                = C.CURLOPT_MAIL_FROM
    CURLOPT_MAIL_RCPT                = C.CURLOPT_MAIL_RCPT
    CURLOPT_FTP_USE_PRET             = C.CURLOPT_FTP_USE_PRET
    CURLOPT_RTSP_REQUEST             = C.CURLOPT_RTSP_REQUEST
    CURLOPT_RTSP_SESSION_ID          = C.CURLOPT_RTSP_SESSION_ID
    CURLOPT_RTSP_STREAM_URI          = C.CURLOPT_RTSP_STREAM_URI
    CURLOPT_RTSP_TRANSPORT           = C.CURLOPT_RTSP_TRANSPORT
    CURLOPT_RTSP_CLIENT_CSEQ         = C.CURLOPT_RTSP_CLIENT_CSEQ
    CURLOPT_RTSP_SERVER_CSEQ         = C.CURLOPT_RTSP_SERVER_CSEQ
    CURLOPT_INTERLEAVEDATA           = C.CURLOPT_INTERLEAVEDATA
    CURLOPT_INTERLEAVEFUNCTION       = C.CURLOPT_INTERLEAVEFUNCTION
    CURLOPT_WILDCARDMATCH            = C.CURLOPT_WILDCARDMATCH
    CURLOPT_CHUNK_BGN_FUNCTION       = C.CURLOPT_CHUNK_BGN_FUNCTION
    CURLOPT_CHUNK_END_FUNCTION       = C.CURLOPT_CHUNK_END_FUNCTION
    CURLOPT_FNMATCH_FUNCTION         = C.CURLOPT_FNMATCH_FUNCTION
    CURLOPT_CHUNK_DATA               = C.CURLOPT_CHUNK_DATA
    CURLOPT_FNMATCH_DATA             = C.CURLOPT_FNMATCH_DATA
    CURLOPT_RESOLVE                  = C.CURLOPT_RESOLVE
    CURLOPT_TLSAUTH_USERNAME         = C.CURLOPT_TLSAUTH_USERNAME
    CURLOPT_TLSAUTH_PASSWORD         = C.CURLOPT_TLSAUTH_PASSWORD
    CURLOPT_TLSAUTH_TYPE             = C.CURLOPT_TLSAUTH_TYPE
    CURLOPT_TRANSFER_ENCODING        = C.CURLOPT_TRANSFER_ENCODING
    CURLOPT_CLOSESOCKETFUNCTION      = C.CURLOPT_CLOSESOCKETFUNCTION
    CURLOPT_CLOSESOCKETDATA          = C.CURLOPT_CLOSESOCKETDATA
    CURLOPT_GSSAPI_DELEGATION        = C.CURLOPT_GSSAPI_DELEGATION
    CURLOPT_DNS_SERVERS              = C.CURLOPT_DNS_SERVERS
    CURLOPT_ACCEPTTIMEOUT_MS         = C.CURLOPT_ACCEPTTIMEOUT_MS
    CURLOPT_TCP_KEEPALIVE            = C.CURLOPT_TCP_KEEPALIVE
    CURLOPT_TCP_KEEPIDLE             = C.CURLOPT_TCP_KEEPIDLE
    CURLOPT_TCP_KEEPINTVL            = C.CURLOPT_TCP_KEEPINTVL
    CURLOPT_SSL_OPTIONS              = C.CURLOPT_SSL_OPTIONS
    CURLOPT_MAIL_AUTH                = C.CURLOPT_MAIL_AUTH
    CURLOPT_SASL_IR                  = C.CURLOPT_SASL_IR
    CURLOPT_XFERINFOFUNCTION         = C.CURLOPT_XFERINFOFUNCTION
    CURLOPT_XOAUTH2_BEARER           = C.CURLOPT_XOAUTH2_BEARER
    CURLOPT_DNS_INTERFACE            = C.CURLOPT_DNS_INTERFACE
    CURLOPT_DNS_LOCAL_IP4            = C.CURLOPT_DNS_LOCAL_IP4
    CURLOPT_DNS_LOCAL_IP6            = C.CURLOPT_DNS_LOCAL_IP6
    CURLOPT_LOGIN_OPTIONS            = C.CURLOPT_LOGIN_OPTIONS
    CURLOPT_SSL_ENABLE_NPN           = C.CURLOPT_SSL_ENABLE_NPN
    CURLOPT_SSL_ENABLE_ALPN          = C.CURLOPT_SSL_ENABLE_ALPN
    CURLOPT_EXPECT_100_TIMEOUT_MS    = C.CURLOPT_EXPECT_100_TIMEOUT_MS
    CURLOPT_PROXYHEADER              = C.CURLOPT_PROXYHEADER
    CURLOPT_HEADEROPT                = C.CURLOPT_HEADEROPT
    CURLOPT_PINNEDPUBLICKEY          = C.CURLOPT_PINNEDPUBLICKEY
    CURLOPT_UNIX_SOCKET_PATH         = C.CURLOPT_UNIX_SOCKET_PATH
    CURLOPT_SSL_VERIFYSTATUS         = C.CURLOPT_SSL_VERIFYSTATUS
    CURLOPT_SSL_FALSESTART           = C.CURLOPT_SSL_FALSESTART
    CURLOPT_PATH_AS_IS               = C.CURLOPT_PATH_AS_IS
    CURLOPT_PROXY_SERVICE_NAME       = C.CURLOPT_PROXY_SERVICE_NAME
    CURLOPT_SERVICE_NAME             = C.CURLOPT_SERVICE_NAME
    CURLOPT_PIPEWAIT                 = C.CURLOPT_PIPEWAIT
)

// CURLE_XXX
const (
    CURLE_OK                       = C.CURLE_OK
    CURLE_UNSUPPORTED_PROTOCOL     = C.CURLE_UNSUPPORTED_PROTOCOL
    CURLE_FAILED_INIT              = C.CURLE_FAILED_INIT
    CURLE_URL_MALFORMAT            = C.CURLE_URL_MALFORMAT
    CURLE_NOT_BUILT_IN             = C.CURLE_NOT_BUILT_IN
    CURLE_COULDNT_RESOLVE_PROXY    = C.CURLE_COULDNT_RESOLVE_PROXY
    CURLE_COULDNT_RESOLVE_HOST     = C.CURLE_COULDNT_RESOLVE_HOST
    CURLE_COULDNT_CONNECT          = C.CURLE_COULDNT_CONNECT
    CURLE_FTP_WEIRD_SERVER_REPLY   = C.CURLE_FTP_WEIRD_SERVER_REPLY
    CURLE_REMOTE_ACCESS_DENIED     = C.CURLE_REMOTE_ACCESS_DENIED
    CURLE_FTP_ACCEPT_FAILED        = C.CURLE_FTP_ACCEPT_FAILED
    CURLE_FTP_WEIRD_PASS_REPLY     = C.CURLE_FTP_WEIRD_PASS_REPLY
    CURLE_FTP_ACCEPT_TIMEOUT       = C.CURLE_FTP_ACCEPT_TIMEOUT
    CURLE_FTP_WEIRD_PASV_REPLY     = C.CURLE_FTP_WEIRD_PASV_REPLY
    CURLE_FTP_CANT_GET_HOST        = C.CURLE_FTP_CANT_GET_HOST
    CURLE_FTP_COULDNT_SET_TYPE     = C.CURLE_FTP_COULDNT_SET_TYPE
    CURLE_PARTIAL_FILE             = C.CURLE_PARTIAL_FILE
    CURLE_FTP_COULDNT_RETR_FILE    = C.CURLE_FTP_COULDNT_RETR_FILE
    CURLE_QUOTE_ERROR              = C.CURLE_QUOTE_ERROR
    CURLE_HTTP_RETURNED_ERROR      = C.CURLE_HTTP_RETURNED_ERROR
    CURLE_WRITE_ERROR              = C.CURLE_WRITE_ERROR
    CURLE_UPLOAD_FAILED            = C.CURLE_UPLOAD_FAILED
    CURLE_READ_ERROR               = C.CURLE_READ_ERROR
    CURLE_OUT_OF_MEMORY            = C.CURLE_OUT_OF_MEMORY
    CURLE_OPERATION_TIMEDOUT       = C.CURLE_OPERATION_TIMEDOUT
    CURLE_FTP_PORT_FAILED          = C.CURLE_FTP_PORT_FAILED
    CURLE_FTP_COULDNT_USE_REST     = C.CURLE_FTP_COULDNT_USE_REST
    CURLE_RANGE_ERROR              = C.CURLE_RANGE_ERROR
    CURLE_HTTP_POST_ERROR          = C.CURLE_HTTP_POST_ERROR
    CURLE_SSL_CONNECT_ERROR        = C.CURLE_SSL_CONNECT_ERROR
    CURLE_BAD_DOWNLOAD_RESUME      = C.CURLE_BAD_DOWNLOAD_RESUME
    CURLE_FILE_COULDNT_READ_FILE   = C.CURLE_FILE_COULDNT_READ_FILE
    CURLE_LDAP_CANNOT_BIND         = C.CURLE_LDAP_CANNOT_BIND
    CURLE_LDAP_SEARCH_FAILED       = C.CURLE_LDAP_SEARCH_FAILED
    CURLE_FUNCTION_NOT_FOUND       = C.CURLE_FUNCTION_NOT_FOUND
    CURLE_ABORTED_BY_CALLBACK      = C.CURLE_ABORTED_BY_CALLBACK
    CURLE_BAD_FUNCTION_ARGUMENT    = C.CURLE_BAD_FUNCTION_ARGUMENT
    CURLE_INTERFACE_FAILED         = C.CURLE_INTERFACE_FAILED
    CURLE_TOO_MANY_REDIRECTS       = C.CURLE_TOO_MANY_REDIRECTS
    CURLE_UNKNOWN_OPTION           = C.CURLE_UNKNOWN_OPTION
    CURLE_TELNET_OPTION_SYNTAX     = C.CURLE_TELNET_OPTION_SYNTAX
    CURLE_PEER_FAILED_VERIFICATION = C.CURLE_PEER_FAILED_VERIFICATION
    CURLE_GOT_NOTHING              = C.CURLE_GOT_NOTHING
    CURLE_SSL_ENGINE_NOTFOUND      = C.CURLE_SSL_ENGINE_NOTFOUND
    CURLE_SSL_ENGINE_SETFAILED     = C.CURLE_SSL_ENGINE_SETFAILED
    CURLE_SEND_ERROR               = C.CURLE_SEND_ERROR
    CURLE_RECV_ERROR               = C.CURLE_RECV_ERROR
    CURLE_SSL_CERTPROBLEM          = C.CURLE_SSL_CERTPROBLEM
    CURLE_SSL_CIPHER               = C.CURLE_SSL_CIPHER
    CURLE_SSL_CACERT               = C.CURLE_SSL_CACERT
    CURLE_BAD_CONTENT_ENCODING     = C.CURLE_BAD_CONTENT_ENCODING
    CURLE_LDAP_INVALID_URL         = C.CURLE_LDAP_INVALID_URL
    CURLE_FILESIZE_EXCEEDED        = C.CURLE_FILESIZE_EXCEEDED
    CURLE_USE_SSL_FAILED           = C.CURLE_USE_SSL_FAILED
    CURLE_SEND_FAIL_REWIND         = C.CURLE_SEND_FAIL_REWIND
    CURLE_SSL_ENGINE_INITFAILED    = C.CURLE_SSL_ENGINE_INITFAILED
    CURLE_LOGIN_DENIED             = C.CURLE_LOGIN_DENIED
    CURLE_TFTP_NOTFOUND            = C.CURLE_TFTP_NOTFOUND
    CURLE_TFTP_PERM                = C.CURLE_TFTP_PERM
    CURLE_REMOTE_DISK_FULL         = C.CURLE_REMOTE_DISK_FULL
    CURLE_TFTP_ILLEGAL             = C.CURLE_TFTP_ILLEGAL
    CURLE_TFTP_UNKNOWNID           = C.CURLE_TFTP_UNKNOWNID
    CURLE_REMOTE_FILE_EXISTS       = C.CURLE_REMOTE_FILE_EXISTS
    CURLE_TFTP_NOSUCHUSER          = C.CURLE_TFTP_NOSUCHUSER
    CURLE_CONV_FAILED              = C.CURLE_CONV_FAILED
    CURLE_CONV_REQD                = C.CURLE_CONV_REQD
    CURLE_SSL_CACERT_BADFILE       = C.CURLE_SSL_CACERT_BADFILE
    CURLE_REMOTE_FILE_NOT_FOUND    = C.CURLE_REMOTE_FILE_NOT_FOUND
    CURLE_SSH                      = C.CURLE_SSH
    CURLE_SSL_SHUTDOWN_FAILED      = C.CURLE_SSL_SHUTDOWN_FAILED
    CURLE_AGAIN                    = C.CURLE_AGAIN
    CURLE_SSL_CRL_BADFILE          = C.CURLE_SSL_CRL_BADFILE
    CURLE_SSL_ISSUER_ERROR         = C.CURLE_SSL_ISSUER_ERROR
    CURLE_FTP_PRET_FAILED          = C.CURLE_FTP_PRET_FAILED
    CURLE_RTSP_CSEQ_ERROR          = C.CURLE_RTSP_CSEQ_ERROR
    CURLE_RTSP_SESSION_ERROR       = C.CURLE_RTSP_SESSION_ERROR
    CURLE_FTP_BAD_FILE_LIST        = C.CURLE_FTP_BAD_FILE_LIST
    CURLE_CHUNK_FAILED             = C.CURLE_CHUNK_FAILED
    CURLE_NO_CONNECTION_AVAILABLE  = C.CURLE_NO_CONNECTION_AVAILABLE
    CURLE_SSL_PINNEDPUBKEYNOTMATCH = C.CURLE_SSL_PINNEDPUBKEYNOTMATCH
    CURLE_SSL_INVALIDCERTSTATUS    = C.CURLE_SSL_INVALIDCERTSTATUS
)