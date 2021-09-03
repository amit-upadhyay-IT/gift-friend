package util

func InHTMLTag(text string) string {
	return `
<HTML>
`+ text+ `
</HTML>
`
}

func InHTMLLinkTag(link string) string {
	return `
<HTML>
<head>
<a href="`+link+`">Click here</a> to you get amazon gift voucher
</head>
</HTML>
`
}

func GetEnterNameHTML() string {
	return `
<html>
<body>

<form action="/gift">
  <label for="firstname">first name please: </label>
  <input type="text" id="firstname" name="firstname"><br>
  <input type="submit" value="Submit">
</form>

</body>
</html>
`
}