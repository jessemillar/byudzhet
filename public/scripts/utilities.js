function trailingZero(amount) { // Add a trailing zero to values that need it
    if (String(amount).match(/\d+\.\d\b/)) {
        return String(amount) + "0";
    }

    return amount;
}

function logout() {
    document.cookie = "id_token=;expires=Thu, 01 Jan 1970 00:00:01 GMT;"; // Delete the cookie by making it expire
    window.location.href = "/";
}

function setCurrentTab(tab) {
    currentTab = tab;

    if (currentTab == "buckets") {
        bootpage.switch("buckets-page");
        setActiveNavigation("buckets-navigation-icon");
    } else if (currentTab == "makeBucket") {
        setActiveNavigation("buckets-navigation-icon");
        document.getElementById("name").focus();
    } else if (currentTab == "expenses") {
        setActiveNavigation("expenses-navigation-icon");
    } else if (currentTab == "logExpense") {
        setActiveNavigation("expenses-navigation-icon");
        document.getElementById("amount").focus();
    } else if (currentTab == "income") {
        setActiveNavigation("income-navigation-icon");
    } else if (currentTab == "logIncome") {
        setActiveNavigation("income-navigation-icon");
        document.getElementById("payer").focus();
    } else if (currentTab == "settings") {
        setActiveNavigation("settings-navigation-icon");
    }
}

function setActiveNavigation(button) {
    // Reset all buttons
    document.getElementById("buckets-navigation-icon").className = "navigation-button";
    document.getElementById("expenses-navigation-icon").className = "navigation-button";
    document.getElementById("income-navigation-icon").className = "navigation-button";
    document.getElementById("history-navigation-icon").className = "navigation-button";
    document.getElementById("settings-navigation-icon").className = "navigation-button";

    // Make the button we care about active
    document.getElementById(button).className += " active";
}

function load() {
    setCurrentTab("buckets");

    getBuckets(populateBuckets);
    getBuckets(populateBucketsDropdown);
    getProjectedIncome(populateProjectedIncome);
    getIncome(populateIncome);
}

function doneLoading(page) {
    $("#" + page + "-page > .loader").hide();
    $("#" + page + "-page > .after-load").show(); // Hide things that shouldn't display until after loading is complete
}
