App.TreetableController = Ember.ObjectController.extend({
    needs: ['treetable'],   // We need this because we are registering an instance of treetable which contains our properties but is not the target of events. See toggleExpand - it needs controllers.treetable. Revisit later

    allExpanded: false,

    toggleAllExpanded: function() {
        Ember.set(this.get('controllers.treetable'), 'allExpanded', !this.get('controllers.treetable.allExpanded'));
    }
});