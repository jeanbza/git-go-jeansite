window.App = Ember.Application.create({
    rootElement: '.jdk-js-ember-widget-demo'
});

// MODEL

App.Widgets = Ember.ArrayController.create({
    items: []
});

// VIEW

App.CostCenter = Ember.Object.extend({
    title: "default",
    content: null
});