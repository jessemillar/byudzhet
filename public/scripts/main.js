var allBuckets; // An array for searching through buckets
var selectedBucket; // A global for keeping track of which bucket is selected in #bucket-dropdown
var projected = false;
var currentTab = "";

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
