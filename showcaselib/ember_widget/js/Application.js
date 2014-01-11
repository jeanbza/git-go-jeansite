App.ApplicationController = Ember.Controller.extend({
    init: function() {
        var self = this;
        
        for(var x = 0; x < 5; x++) {
            var tempView = App.Widget.create({
                data: Ember.makeArray([
                    App.CostCenter.create({title: "Group A", content: self.getRandomArr(8)}),
                    App.CostCenter.create({title: "Group B", content: self.getRandomArr(8)}),
                    App.CostCenter.create({title: "Group C", content: self.getRandomArr(8)})
                ]),
            });
            App.Widgets.items.pushObject(tempView);
        }
    },

    getRandomArr: function(size) {
        var randomArr = [];

        for(var x = 0; x < size; x++) {
            randomArr.push(this.getRandomInt(0, 10));
        }

        return randomArr;
    },

    getRandomInt: function(min, max) {
        return Math.floor(Math.random() * (max - min + 1)) + min;
    },

    switchMode: function(view) {
        view.switchMode();
    }
});