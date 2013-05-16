var Inventory = {
	PKWhitelist: /^[\d*+-\/()]+$/,
	TimeFmt: /^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}$/,
	init: function() {
		rows = [
			$("#InventoryForm .resonators input"),
			$("#InventoryForm .xmps input"),
			$("#InventoryForm .powercubes input"),
			$("#InventoryForm .mods input")
		]
		for (var i = 0; i < rows.length; i++) {
			rows[i].each(Inventory.copyValue).change(Inventory.updateTotal).first().trigger("change")
		}
		$("#PortalKeyCalc").keyup(function() {
			console.log($(this).val())
			Inventory.calculatePKs($(this), $("input[name=PortalKeys]"), $("#PortalKeyDisp"))
		})
		$("#InventoryForm").submit(Inventory.submitListener)
	},
	toRFC3339: function(d) {
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
	},
	setError: function(e, err) {
		var group = e.parents(".control-group")
		var errNode = group.find(".err")
		if (errNode.length == 0) {
			errNode = $('<span class="err help-inline"></span>')
			errNode.insertAfter(e)
		}
		errNode.text(err)
		group.addClass("error")
	},
	clearError: function(e) {
		var group = e.parents(".control-group")
		group.find(".err").text("")
		group.removeClass("error")
	},
	copyValue: function() {
		$(this).data("val", $(this).val())
	},
	validateTime: function(from, to) {
		var v = from.val()
		if (!Inventory.TimeFmt.test(v)) {
			Inventory.setError(from, "Invalid format. Expect: yyyy-mm-ddThh:mm")
			return false
		}
		to.val(Inventory.toRFC3339(v))
		return true
	},
	calculatePKs: function(from, to, disp) {
		var v = from.val()
		if (v == "") {
			disp.text("0")
			return true
		}
		if (!Inventory.PKWhitelist.test(v)) {
			Inventory.setError(from, "Invalid character(s). Only numbers and +, -, *, /, (, and ) allowed.")
			return false
		}
		Inventory.clearError(from)

		var result
		try {
			eval("result = "+v)
		} catch (e) {
			if (e instanceof SyntaxError) {
				// Inventory.setError(from, e.message)
				return false
			} else {
				console.log(e.message)
			}
		}
		to.val(result)
		disp.text(result)
		return true
	},
	rowTotal: function(row) {
		var total = 0
		row.find("input").each(function() {
			total += parseInt($(this).data("val"))
		})
		row.find(".row-total").text(total)
	},
	submitListener: function() {
		var self = $(this)
		switch (false) {
		case Inventory.validateTime(self.find("#EntryTime"), self.find("input[name=Time]")):
			return false
		}
		return true
	},
	updateTotal: function() {
		$(this).data("val", $(this).val())
		Inventory.rowTotal($(this).parents(".control-group"))
	}
}

$(Inventory.init)
