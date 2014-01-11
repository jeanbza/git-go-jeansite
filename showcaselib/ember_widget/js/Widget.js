App.Widget = Ember.View.extend({
    templateName: "widget",
    someArr: [1,5,3],
    title: "Some Widget",
    graphElem: ".graphHere",
    graphMode: true,

    didInsertElement: function() {
        this.graphData();
    },

    switchMode: function() {
        this.set("graphMode", !this.graphMode);
        if(this.graphMode) {
            this.graphData();
        } else {
            this.$().find(this.graphElem).empty();
        }
    },

    graphData: function() {
        var self = this;
        var data = [];
        var graphClass = "graphArea";

        $(document).ready(function() {
            self.$().find(self.graphElem).append("<div class='"+graphClass+"'></div>");

            $(self.data).each(function(){
                var datum = {name: this.title, data: this.content};
                data.push(datum);
            });

            self.$().find("."+graphClass).highcharts({
                chart: {
                    height: 250,
                    marginRight: 10
                },
                title: {
                    text: 'Scores',
                    x: -20
                },
                xAxis: {
                    categories: ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun', 'Jul', 'Aug'],
                    labels: {
                        rotation: -45
                    }
                },
                yAxis: {
                    title: {
                        text: 'Score'
                    },
                    plotLines: [{
                        value: 0,
                        width: 1,
                        color: '#808080'
                    }]
                },
                legend: {
                    layout: 'vertical',
                    align: 'right',
                    verticalAlign: 'middle',
                    borderWidth: 0
                },
                series: data,
                credits: {
                    enabled: false
                }
            });
        });
    }
});