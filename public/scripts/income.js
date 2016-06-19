function logIncome() {
    body = {
        amount: $("#log-income-amount").val(),
        payer: $("#log-income-payer").val(),
    };

    $.ajax("/api/income", {
        "data": JSON.stringify(body),
        "type": "POST",
        "processData": false,
        "contentType": "application/json",
        "success": function(data) {
            window.location.href = "/frontend#income";
            location.reload();
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

function setProjectedIncome() {
    body = {
        amount: $("#amount").val()
    };

    var httpType;

    if (projected) {
        httpType = "PUT";
    } else {
        httpType = "POST";
    }

    $.ajax("/api/projected", {
        "data": JSON.stringify(body),
        "type": httpType,
        "processData": false,
        "contentType": "application/json",
        "success": function(data) {
            window.location.href = "/frontend#settings";
            location.reload();
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
        amountSpan.appendChild(document.createTextNode("$" + trailingZero(income[i].amount)));

        amount.appendChild(amountSpan);

        row.appendChild(payer);
        row.appendChild(amount);
        li.appendChild(row);

        document.getElementById("income-list").appendChild(li);
    }

    doneLoading("income");
}

function getProjectedIncome(callback) {
    $.get("/api/projected", function(data) {
        if (data.amount > 0) {
            projected = true;

            if (callback) {
                callback(data);
            } else {
                return data;
            }
        }
    });
}

function populateProjectedIncome(data) {
    // Set projected income value on expenses tab
    var expensesProgress = document.createElement("div");

    if (Number(data.spent) < Number(data.amount * 0.5)) {
        expensesProgress.className = "progress-bar progress-bar-success";
    } else if (Number(data.spent) < Number(data.amount * 0.75)) {
        expensesProgress.className = "progress-bar progress-bar-warning";
    } else {
        expensesProgress.className = "progress-bar progress-bar-danger";
    }

    expensesProgress.style.width = data.spent / data.amount * 100 + "%"; // Populate this with a calculated value

    $("#expenses-progress").append(expensesProgress);

    $("#expenses-ratio").text("$" + trailingZero(data.spent) + " / $" + trailingZero(data.amount));

    // Set projected income value on income tab
    var incomeProgress = document.createElement("div");

    if (Number(data.earned) < Number(data.amount * 0.5)) {
        incomeProgress.className = "progress-bar progress-bar-danger";
    } else if (Number(data.earned) < Number(data.amount * 0.75)) {
        incomeProgress.className = "progress-bar progress-bar-warning";
    } else {
        incomeProgress.className = "progress-bar progress-bar-success";
    }

    incomeProgress.style.width = data.earned / data.amount * 100 + "%"; // Populate this with a calculated value

    $("#income-progress").append(incomeProgress);

    $("#income-ratio").text("$" + trailingZero(data.earned) + " / $" + trailingZero(data.amount));

    // Set projected income value on settings tab
    $("#amount").val(trailingZero(data.amount));
}
