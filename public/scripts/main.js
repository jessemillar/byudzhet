// Make the app work as a single-page app on iOS devices
(function(a, b, c) {
    if (c in b && b[c]) {
        var d, e = a.location,
            f = /^(a|html)$/i;
        a.addEventListener("click", function(a) {
            d = a.target;
            while (!f.test(d.nodeName)) d = d.parentNode;
            "href" in d && (chref = d.href).replace(e.href, "").indexOf("#") && (!/^[a-z\+\.\-]+:/i.test(chref) || chref.indexOf(e.protocol + "//" + e.host) === 0) && (a.preventDefault(), e.href = d.href)
        }, !1)
    }
})(document, window.navigator, "standalone");

function init() {
    page = window.location.pathname;

    if (page == "/bucket") {
        setActiveNavigation("bucket");
    } else if (page == "/expense") {
        setActiveNavigation("expense");
    } else if (page == "/income") {
        setActiveNavigation("income");
    } else if (page == "/setting") {
        setActiveNavigation("setting");
    } else if (page == "/log") {
        setActiveNavigation("expense");

        document.getElementById("amount").focus();
    }
}

function setActiveNavigation(button) {
    // Reset all buttons
    document.getElementById("bucket").className = "navigation-button";
    document.getElementById("expense").className = "navigation-button";
    document.getElementById("income").className = "navigation-button";
    document.getElementById("setting").className = "navigation-button";

    // Make the button we care about active
    document.getElementById(button).className += " active";
}

function logExpense() {
    amount = $("#amount").val()
    recipient = $("#recipient").val()
    note = $("#note").val()

    body = {
        amount: amount,
        recipient: recipient,
        note: note
    };

    console.log(JSON.stringify(body));

    $.post("/expense", JSON.stringify(body), function(data) {
        console.log(data);
    });
}
