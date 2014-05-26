App = Ember.Application.create({
    rootElement: ".jdk-js-ember-treetable-demo"
});

App.ApplicationController = Ember.Controller.extend({
    expandRandomItem: function() {
        App.TreeNodeController.controllerForNodeById(10).bubbleExpanded();
    }
});