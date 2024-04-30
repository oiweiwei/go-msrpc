# MS Kerberos Support

# SHA-2 Support

[Link](https://lists.fedorahosted.org/archives/list/freeipa-users@lists.fedorahosted.org/thread/FITCJOXX2QQ4HEXEK4PDJWFZJ2C33FAZ/)

```
Windows systems do not support RFC8009 encryption types (SHA-2-based) at
all. Microsoft engineer gave a hint they are working on their support
for ~2025 but it will not be backported.

IPA KDC defaults to RFC8009 and only falls back to SHA-1-based ones
for trust to Active Directory. There are few places in MIT Kerberos KDC
where a choice of the signature or encryption type is made based on the
strongest key available for the krbtgt/... principal, which is always a
SHA-2-based one for new IPA deployments. For cross-realm operations we
have special logic to fall back to SHA-1-based ones for AD DCs. For
in-realm operations we don't and shouldn't as it would be a security
issue (downgrade of encryption algorithm to a less secure one).
```
