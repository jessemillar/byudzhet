var allBuckets; // An array for searching through buckets
var selectedBucket; // A global for keeping track of which bucket is selected in #bucket-dropdown
var projected = false;
var currentTab = "";

(function(a, b, c) { // Make the app work as a single-page app on iOS devices
    if (c in b && b[c]) {
        var d, e = a.location,
            f = /^(a|html)$/i;
        a.addEventListener("click", function(a) {
            d = a.target;
            while (!f.test(d.nodeName)) d = d.parentNode;
            "href" in d && (chref = d.href).replace(e.href, "").indexOf("#") && (!/^[a-z\+\.\-]+:/i.test(chref) || chref.indexOf(e.protocol + "//" + e.host) === 0) && (a.preventDefault(), e.href = d.href)
        }, !1);
    }
})(document, window.navigator, "standalone");

$(function() {
    $("body").on('click', '.dropdown-menu li a', function() { // Populate #bucket-dropdown with selected item
        selectedBucket = $(this).text();
        $("#bucket-dropdown").html($(this).text() + " <span class='caret'></span>");
    });

    $("#make-bucket-amount").on("input", function(e) { // Make the input auto-insert the decimal
        $("#make-bucket-amount").val($("#make-bucket-amount").val().replace(/[^\d]/g, '').replace(/(\d\d?)$/, '.$1'));
    });

    $("#log-income-amount").on("input", function(e) { // Make the input auto-insert the decimal
        $("#log-income-amount").val($("#log-income-amount").val().replace(/[^\d]/g, '').replace(/(\d\d?)$/, '.$1'));
    });

    $("#log-expense-amount").on("input", function(e) { // Make the input auto-insert the decimal
        $("#log-expense-amount").val($("#log-expense-amount").val().replace(/[^\d]/g, '').replace(/(\d\d?)$/, '.$1'));
    });
});
