$(function() {
	$("#InventoryForm input[name=Timezone]").val()
	$("#InventoryForm").submit(function() {
		var self = $(this)
		switch (false) {
		case validateTime(self.find("#EntryTime"), self.find("input[name=Time]")):
			return false
		}
		return true
	})
})
var toRFC3339 = function(d) {
	d += ":00"
	var t = new Date()
	var tz = t.getTimezoneOffset()
	if (tz > 0)
		d += "-"
	var h = Math.abs(parseInt(tz/60))
	if (h < 10)
		h = "0"+h
	var m = tz%60
	if (m == 0)
		m = "00"
	d += h+":"+m
	return d
}
var validateTime = function(from, to) {
	var fmt = /^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}$/
	var v = from.val()
	if (!fmt.test(v)) {
		setError(e, "Invalid format. Expect: yyyy-mm-ddThh:mm")
		return false
	}
	to.val(toRFC3339(v))
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
