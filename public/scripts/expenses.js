function logExpense() {
    for (var i in allBuckets) { // Find the ID of the selected bucket
        if (allBuckets[i].name == selectedBucket) {
            selectedBucket = allBuckets[i].id;
        }
    }

    body = {
        bucket: selectedBucket.toString(),
        amount: $("#log-expense-amount").val(),
        recipient: $("#log-expense-recipient").val(),
        note: $("#log-expense-note").val()
    };

    $.ajax("/api/expense", {
        "data": JSON.stringify(body),
        "type": "POST",
        "processData": false,
        "contentType": "application/json",
        "success": function(data) {
            window.location.href = "/frontend#expenses";
            location.reload();
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
        amountSpan.appendChild(document.createTextNode("$" + trailingZero(expenses[i].amount)));
        note.appendChild(document.createTextNode(expenses[i].note));

        amount.appendChild(amountSpan);

        topRow.appendChild(recipient);
        topRow.appendChild(amount);
        bottomRow.appendChild(note);

        li.appendChild(topRow);
        li.appendChild(bottomRow);

        $("#expenses-page #expenses-list").append(li);
    }

    doneLoading("expenses");
}
