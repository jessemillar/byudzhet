var bootpage = {
    currentPage: ""
};

bootpage.switch = function(page) {
    if (self.currentPage) {
        document.getElementById(self.currentPage).style.display = "none"; // Hide the current page
    }

    document.getElementById(page).style.display = "block"; // Hide the current page
    self.currentPage = page;
};
