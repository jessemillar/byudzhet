function trailingZero(amount) {
    if (String(amount).match(/\d+\.\d\b/)) {
        return String(amount) + "0";
    }

    return amount;
}

function setCurrentTab(tab) {
    currentTab = "#" + tab;
}

function load(bootpageCall) {
    page = currentTab; // Get the page we're on

    if (page == "#buckets") {
        if (!bootpageCall) {
            bootpage.switch("buckets-page");
        }

        setActiveNavigation("buckets");

        getBuckets(populateBuckets);
    } else if (page == "#makeBucket") {
        setActiveNavigation("buckets");

        document.getElementById("name").focus();
    } else if (page == "#expenses") {
        setActiveNavigation("expenses");

        getProjectedIncome(populateProjectedIncome);
        getExpenses(populateExpenses);
    } else if (page == "#logExpense") {
        setActiveNavigation("expenses");

        getBuckets(populateBucketsDropdown);
        document.getElementById("amount").focus();
    } else if (page == "#income") {
        setActiveNavigation("income");

        getProjectedIncome(populateProjectedIncome);
        getIncome(populateIncome);
    } else if (page == "#logIncome") {
        setActiveNavigation("income");

        document.getElementById("payer").focus();
    } else if (page == "#settings") {
        getProjectedIncome(populateProjectedIncome);

        setActiveNavigation("settings");
    }
}

function hideLoader(page) {
    console.log("Hiding loader", "." + page + "-page > .loader");

    $("#" + page + "-page > .loader").hide();
    $("#" + page + "-page > .after-load").show(); // Hide things that shouldn't display until after loading is complete
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
