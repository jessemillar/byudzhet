var allBuckets; // An array for searching through buckets
var selectedBucket; // A global for keeping track of which bucket is selected in #bucket-dropdown

(function(a, b, c) { // Make the app work as a single-page app on iOS devices
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

$(function() { // Populate #bucket-dropdown with selected item
    $("body").on('click', '.dropdown-menu li a', function() {
        selectedBucket = $(this).text();
        $("#bucket-dropdown").html($(this).text() + " <span class='caret'></span>");
    });
});

function init() {
    page = window.location.pathname; // Get the page we're on

    if (page == "/buckets") {
        setActiveNavigation("buckets");

        getBuckets(populateBuckets);
    } else if (page == "/buckets/make") {
        setActiveNavigation("buckets");

        document.getElementById("name").focus();
    } else if (page == "/expenses") {
        setActiveNavigation("expenses");

        getExpenses(populateExpenses);
    } else if (page == "/expenses/log") {
        setActiveNavigation("expenses");

        getBuckets(populateBucketsDropdown);
        document.getElementById("amount").focus();
    } else if (page == "/income") {
        setActiveNavigation("income");

        getIncome(populateIncome);
    } else if (page == "/income/log") {
        setActiveNavigation("income");

        document.getElementById("payer").focus();
    } else if (page == "/income") {
        setActiveNavigation("income");
    } else if (page == "/settings") {
        setActiveNavigation("settings");
    }
}

function hideLoader() {
    $("#loader").hide();
}

function logout() {
    document.cookie = "id_token=;expires=Thu, 01 Jan 1970 00:00:01 GMT;"; // Delete the cookie by making it expire
    window.location.href = "/";
}

function setActiveNavigation(button) {
    // Reset all buttons
    document.getElementById("buckets").className = "navigation-button";
    document.getElementById("expenses").className = "navigation-button";
    document.getElementById("income").className = "navigation-button";
    document.getElementById("settings").className = "navigation-button";

    // Make the button we care about active
    document.getElementById(button).className += " active";
}

function logExpense() {
    for (var i in allBuckets) { // Find the ID of the selected bucket
        if (allBuckets[i].name == selectedBucket) {
            selectedBucket = allBuckets[i].id;
        }
    }

    console.log(selectedBucket, allBuckets);

    body = {
        bucket: selectedBucket.toString(),
        amount: $("#amount").val(),
        recipient: $("#recipient").val(),
        note: $("#note").val()
    };

    console.log(body);

    $.ajax("/api/expense", {
        "data": JSON.stringify(body),
        "type": "POST",
        "processData": false,
        "contentType": "application/json",
        "success": function(data) {
            window.location.href = "/expenses";
        }
    });
}

function getExpenses(callback) {
    $.get("/api/expense", function(data) {
        if (callback) {
            callback(data);
        } else {
            return data;
        }
    });
}

function populateExpenses(expenses) {
    for (var i in expenses) {
        var li = document.createElement("li");
        var topRow = document.createElement("div");
        var bottomRow = document.createElement("div");
        var recipient = document.createElement("div");
        var note = document.createElement("div");
        var amount = document.createElement("div");
        var amountSpan = document.createElement("span");

        li.className = "list-group-item";
        topRow.className = "row list-row";
        recipient.className = "col-xs-7 list-title";
        amount.className = "col-xs-5 badge-amount";
        amountSpan.className = "badge";
        bottomRow.className = "row list-row";
        note.className = "col-xs-12";

        recipient.appendChild(document.createTextNode(expenses[i].recipient));
        amountSpan.appendChild(document.createTextNode("$" + expenses[i].amount));
        note.appendChild(document.createTextNode(expenses[i].note));

        amount.appendChild(amountSpan);

        topRow.appendChild(recipient);
        topRow.appendChild(amount);
        bottomRow.appendChild(note);

        li.appendChild(topRow);
        li.appendChild(bottomRow);

        document.getElementById("expenses-list").appendChild(li);
    }

    hideLoader();
}

function makeBucket() {
    body = {
        amount: $("#amount").val(),
        name: $("#name").val()
    };

    $.ajax("/api/bucket", {
        "data": JSON.stringify(body),
        "type": "POST",
        "processData": false,
        "contentType": "application/json",
        "success": function(data) {
            window.location.href = "/buckets";
        }
    });
}

function getBuckets(callback) {
    $.get("/api/bucket", function(data) {
        if (callback) {
            callback(data);
        } else {
            return data;
        }
    });
}

function getBucketByName(bucket, callback) {
    $.get("/api/bucket/" + bucket.name, function(data) {
        if (callback) {
            callback(data);
        } else {
            return data;
        }
    });
}

function populateBuckets(buckets) {
    for (var i in buckets) {
        var li = document.createElement("li");
        var topRow = document.createElement("div");
        var bottomRow = document.createElement("div");
        var name = document.createElement("div");
        var ratio = document.createElement("div");
        var ratioSpan = document.createElement("span");
        var progressCol = document.createElement("div");
        var progressWrapper = document.createElement("div");
        var progress = document.createElement("div");

        li.className = "list-group-item";
        topRow.className = "row list-row";
        name.className = "col-xs-7 list-title";
        ratio.className = "col-xs-5 badge-amount";
        ratioSpan.className = "badge";
        bottomRow.className = "row list-row";
        progressCol.className = "col-xs-12";
        progressWrapper.className = "progress";

        if (buckets[i].spent > buckets[i].amount) {
            progress.className = "progress-bar progress-bar-danger";
        } else if (buckets[i].spent > buckets[i].amount * 0.75) {
            progress.className = "progress-bar progress-bar-warning";
        } else {
            progress.className = "progress-bar progress-bar-success";
        }

        progress.style.width = buckets[i].spent / buckets[i].amount * 100 + "%"; // Populate this with a calculated value

        name.appendChild(document.createTextNode(buckets[i].name));
        ratioSpan.appendChild(document.createTextNode("$" + buckets[i].spent + " / " + "$" + buckets[i].amount));

        ratio.appendChild(ratioSpan);

        progressWrapper.appendChild(progress);
        progressCol.appendChild(progressWrapper);

        topRow.appendChild(name);
        topRow.appendChild(ratio);
        bottomRow.appendChild(progressCol);

        li.appendChild(topRow);
        li.appendChild(bottomRow);

        document.getElementById("buckets-list").appendChild(li);
    }

    hideLoader();
}

function populateBucketsDropdown(buckets) {
    allBuckets = buckets;

    for (var i in buckets) {
        var li = document.createElement("li");
        var a = document.createElement("a");

        a.href = "#";

        a.appendChild(document.createTextNode(buckets[i].name));
        li.appendChild(a);

        document.getElementById("bucket-dropdown-options").appendChild(li);

        $('#bucket-dropdown-options').trigger("chosen:updated");
    }
}

function logIncome() {
    body = {
        amount: $("#amount").val(),
        payer: $("#payer").val(),
    };

    $.ajax("/api/income", {
        "data": JSON.stringify(body),
        "type": "POST",
        "processData": false,
        "contentType": "application/json",
        "success": function(data) {
            window.location.href = "/income";
        }
    });
}

function getIncome(callback) {
    $.get("/api/income", function(data) {
        if (callback) {
            callback(data);
        } else {
            return data;
        }
    });
}

function getIncomeTotal(bucketID, callback) {
    getIncome(function(income) {
        var total = 0;

        for (var i in income) {
            total += parseInt(expenses[i].amount);
        }

        if (callback) {
            callback(total);
        } else {
            return total;
        }
    });
}

function populateIncome(income) {
    for (var i in income) {
        var li = document.createElement("li");
        var row = document.createElement("div");
        var payer = document.createElement("div");
        var amount = document.createElement("div");
        var amountSpan = document.createElement("span");

        li.className = "list-group-item";
        row.className = "row list-row";
        payer.className = "col-xs-8";
        amount.className = "col-xs-4 badge-amount";
        amountSpan.className = "badge";

        payer.appendChild(document.createTextNode(income[i].payer));
        amountSpan.appendChild(document.createTextNode("$" + income[i].amount));

        amount.appendChild(amountSpan);

        row.appendChild(payer);
        row.appendChild(amount);
        li.appendChild(row);

        document.getElementById("income-list").appendChild(li);
    }

    hideLoader();
}

function getUserEmail(id, callback) {
    $.get("/api/user/id/" + id, function(data) {
        if (callback) {
            callback(data);
        } else {
            return data;
        }
    });
}
