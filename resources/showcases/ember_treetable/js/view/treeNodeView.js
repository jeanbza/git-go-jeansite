App.TreeNodeView = Ember.View.extend({
    tagName: 'li',
    templateName: 'tree-node',
    classNames: ['tree-node'],
    classNameBindings: ['childrenAreInvisible'],
    childrenAreInvisible: false,

    allExpandedDidChange: function() {
        this.set('childrenAreInvisible', !this.get('controller.isExpanded'));
    }.observes('controller.isExpanded')
});