"use strict";

const url = "https://github.com/packtci/go-template-example-with-circle-ci"

describe("Web Scraping", function() {
    it("Visits the go-template-example-with-circle-ci repository", function() {
        cy.visit(url);

        cy
        .get(".markdown-body.entry-content h1")
        .should("contain", "Repository Details");
    });
});