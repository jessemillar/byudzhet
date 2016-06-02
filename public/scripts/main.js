var allBuckets; // An array for searching through buckets
var selectedBucket; // A global for keeping track of which bucket is selected in #bucket-dropdown
var projected = false;
var currentTab = "#buckets";

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

$(function() { // Populate #bucket-dropdown with selected item
    $("body").on('click', '.dropdown-menu li a', function() {
        selectedBucket = $(this).text();
        $("#bucket-dropdown").html($(this).text() + " <span class='caret'></span>");
    });

    $("#amount").on("input", function(e) {
        $("#amount").val($("#amount").val().replace(/[^\d]/g, '').replace(/(\d\d?)$/, '.$1'));
    });
});
