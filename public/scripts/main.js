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

    if (page == "/buckets") {
        setActiveNavigation("buckets");

        getBuckets(populateBuckets);
    } else if (page == "/buckets/make") {
        setActiveNavigation("buckets");

        document.getElementById("amount").focus();
    } else if (page == "/expenses") {
        setActiveNavigation("expenses");

        getExpenses(populateExpenses);
    } else if (page == "/expenses/log") {
        setActiveNavigation("expenses");

        document.getElementById("amount").focus();
    } else if (page == "/income") {
        setActiveNavigation("income");
    } else if (page == "/settings") {
        setActiveNavigation("settings");
    }
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
    body = {
        amount: $("#amount").val(),
        recipient: $("#recipient").val(),
        note: $("#note").val()
    };

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

function getExpensesTotal(bucketID, callback) {
    getExpenses(function(expenses) {
        var total = 0;

        for (var i in expenses) {
            if (expenses[i].bucket == bucketID) {
                total += parseInt(expenses[i].amount);
            }
        }

        if (callback) {
            callback(total);
        } else {
            return total;
        }
    });
}

function populateExpenses(expenses) {
    for (var i in expenses) {
        var li = document.createElement("li");
        var row = document.createElement("div");
        var receipient = document.createElement("div");
        var note = document.createElement("div");
        var amount = document.createElement("div");
        var amountSpan = document.createElement("span");

        li.className = "list-group-item";
        row.className = "row list-row";
        receipient.className = "col-xs-3";
        note.className = "col-xs-7";
        amount.className = "col-xs-2";
        amountSpan.className = "badge";

        receipient.appendChild(document.createTextNode(expenses[i].recipient));
        note.appendChild(document.createTextNode(expenses[i].note));
        amountSpan.appendChild(document.createTextNode("$" + expenses[i].amount));

        amount.appendChild(amountSpan);

        row.appendChild(receipient);
        row.appendChild(note);
        row.appendChild(amount);
        li.appendChild(row);

        document.getElementById("expenses-list").appendChild(li);
    }
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

function getBucketID(bucket, callback) {
    $.get("/api/bucket/" + bucket, function(data) {
        if (callback) {
            callback(data.id);
        } else {
            return data;
        }
    });
}

function populateBuckets(buckets) {
    for (var i in buckets) {
        getBucketID(buckets[i].name, function(id) {
            getExpensesTotal(id, function(total) {
                var col = document.createElement("div");
                var name = document.createElement("div");
                var progressCol = document.createElement("div");
                var progressWrapper = document.createElement("div");
                var progress = document.createElement("div");

                col.className = "col-xs-12";
                name.className = "col-xs-4";
                progressCol.className = "col-xs-8";
                progressWrapper.className = "progress";
                progress.className = "progress-bar progress-bar-warning";

                progress.style.width = total / buckets[i].amount * 100 + "%"; // Populate this with a calculated value

                name.appendChild(document.createTextNode(buckets[i].name));

                progressWrapper.appendChild(progress);
                progressCol.appendChild(progressWrapper);

                col.appendChild(name);
                col.appendChild(progressCol);

                document.getElementById("buckets-list").appendChild(col);
            });
        });
    }
}
