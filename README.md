> It is dangerous to be right when the government is wrong.
>    -- Voltaire
   
# What?
`fakid` generates a **fake identity**.

Example:
```
% fakid
First name: Leisa
Last name: Pecor
Nickname: peesa83
Sex:  Female
Country:  USA
Profession:  Nuclear Engineer
Email:  leisa.pecor20@mailinator.com
Passphrase:  [calculator pucker slogan gratitude]
Blood type:  A+
Likes:  [Kitties]
```

# Why?

* For a fiction character.
* To stay anonymous online. Some online services may require a "true" identity, with name, age, etc. `fakid` creates one for you.
* Pen name. To experiment with different styles of writing.
* With an *identity* is easier to build a readership, since your readers can identify with you and learn to know what to expect from your writings/blog, but you may want to stay anonymous when speaking about taboos, your secret life, politics, etc.

# Install

First [install Golang](https://golang.org/doc/install) then do

```
go get gitlab.com/xojoc/fakid
go install gitlab.com/xojoc/fakid
```
and run
```
fakid
```
Note1: the email address is real. You can use it.

Note2: `fakid` generates both passphrases and passwords. **Passphrases** are to be preferred, but if a website imposes stupid rules like at least an Uppercase, a digit, a weird character, then you can always use the password.

Note3: you can hide your name, but be aware that you could be identified by your way of writing. Frequencies of letters/words, length of phrases, etc. all reveal bits of you.

# Privacy
If you care about privacy/anonymity also check out:

* [IceCat](https://www.gnu.org/software/gnuzilla/)
* [Tor](https://www.torproject.org/)
* [Freenet](https://freenetproject.org/)

You can use [Typed.pw](http://typed.pw) to write a one-off article.

# Todo
* Short bio
* Long bio
* Fake social messages (with hashtags, @nicknames, ect.)
* Favorite food
* Likes: extend list
* Hates: extend list
* Profession: a better list
* Countries: extend list (Only USA is used right now)
* Height/weight
* Age/tropical zodiac
* Sport
* Hobby
* Avatar (but how?)

# Who?
Written by <http://xojoc.pw>

# License
Public domain. Because copyright laws are lame.

# [Donate](http://xojoc.pw/donate.html)
