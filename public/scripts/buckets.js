function makeBucket() {
    body = {
        amount: $("#make-bucket-amount").val(),
        name: $("#make-bucket-name").val()
    };

    $.ajax("/api/bucket", {
        "data": JSON.stringify(body),
        "type": "POST",
        "processData": false,
        "contentType": "application/json",
        "success": function(data) {
            window.location.href = "/frontend#buckets";
            location.reload();
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

        if (Number(buckets[i].spent) > Number(buckets[i].amount)) {
            progress.className = "progress-bar progress-bar-danger";
        } else if (Number(buckets[i].spent) > Number(buckets[i].amount * 0.75)) {
            progress.className = "progress-bar progress-bar-warning";
        } else {
            progress.className = "progress-bar progress-bar-success";
        }

        progress.style.width = buckets[i].spent / buckets[i].amount * 100 + "%"; // Populate this with a calculated value

        name.appendChild(document.createTextNode(buckets[i].name));
        ratioSpan.appendChild(document.createTextNode("$" + trailingZero(buckets[i].spent) + " / " + "$" + trailingZero(buckets[i].amount)));

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

    doneLoading("buckets");
}

function populateBucketsDropdown(buckets) {
    allBuckets = buckets;

    for (var i in buckets) {
        var li = document.createElement("li");
        var a = document.createElement("a");

        a.href = "javascript:void(0)"; // Don't reload the page on iOS

        a.appendChild(document.createTextNode(buckets[i].name));
        li.appendChild(a);

        $("#bucket-dropdown-options").append(li);
        $("#bucket-dropdown-options").trigger("chosen:updated");
    }
}
