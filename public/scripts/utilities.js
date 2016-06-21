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
        bootpage.show("buckets-page");
        setActiveNavigation("buckets-navigation-icon");
    } else if (currentTab == "make-bucket") {
        bootpage.show("make-bucket-page");
        setActiveNavigation("buckets-navigation-icon");
        focusMakeBucket();
    } else if (currentTab == "expenses") {
        bootpage.show("expenses-page");
        setActiveNavigation("expenses-navigation-icon");
    } else if (currentTab == "log-expense") {
        bootpage.show("log-expense-page");
        setActiveNavigation("expenses-navigation-icon");
        focusLogExpense();
    } else if (currentTab == "income") {
        bootpage.show("income-page");
        setActiveNavigation("income-navigation-icon");
    } else if (currentTab == "log-income") {
        bootpage.show("log-income-page");
        setActiveNavigation("income-navigation-icon");
        focusLogIncome();
    } else if (currentTab == "settings") {
        bootpage.show("settings-page");
        setActiveNavigation("settings-navigation-icon");
    }
}

function focusMakeBucket() {
    document.getElementById("make-bucket-name").focus();
}

function focusLogExpense() {
    document.getElementById("log-expense-amount").focus();
}

function focusLogIncome() {
    document.getElementById("log-income-payer").focus();
}

function setActiveNavigation(button) {
    // Reset all buttons
    document.getElementById("buckets-navigation-icon").className = "navigation-button";
    document.getElementById("expenses-navigation-icon").className = "navigation-button";
    document.getElementById("income-navigation-icon").className = "navigation-button";
    document.getElementById("settings-navigation-icon").className = "navigation-button";

    // Make the button we care about active
    document.getElementById(button).className += " active";
}

function load() {
    if (window.location.hash) {
        var page = window.location.hash.substr(1);
        setCurrentTab(page);
    } else {
        setCurrentTab("buckets");
    }

    getBuckets(populateBuckets);
    getBuckets(populateBucketsDropdown);
    getExpenses(populateExpenses);
    getIncome(populateIncome);
    getProjectedIncome(populateProjectedIncome);
}

function doneLoading(page) {
    $("#" + page + "-page > .loader").hide();
    $("#" + page + "-page > .after-load").show(); // Hide things that shouldn't display until after loading is complete
}
