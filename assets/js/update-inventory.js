$(function() {
	$("#InventoryForm").submit(function() {
		var self = $(this)
		if (validateTime(self.find("input[name=Time]")))
		return false
	})
})
var validateTime = function(e) {
	var fmt = /^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}$/
	var v = e.val()
	if (!fmt.test(v)) {
		setError(e, "Invalid format. Expect: yyyy-mm-ddThh:mm")
		return false
	}
	
	return true
}
var setError = function(e, err) {
	var sib
	var sibs = e.siblings('.err')
	if (sibs.length == 0) {
		sib = $('<span class="err inline-help"></span>')
		sib.insertAfter(e)
	} else {
		sib = $(sibs[0])
	}
	sib.text(err)
	e.parents(".control-group").addClass("error")
}
/* Some JavaScript for Portal Keys later
var result = 5
try {
	/[\d*+-/()]/
	resut = eval("1+2")
} catch (e) {
	if (e instanceof SyntaxError) {

	} else {
		console.log(e.message)
	}
}
*/
