$(document).ready(function(){
    refreshTreeBindings();
});

/**
 * This function refreshes the bindings on the tree
 * @return {[type]} [description]
 */
function refreshTreeBindings() {
    // Reveal some random node - this is just for demo purposes - remove in production
    $(".js-reveal-random-node").click(function(){
        var numItemsInTree = $(this).attr('data-max-items');
        expandAndCheckTreeNode(getRandomNumber(0, numItemsInTree-1));
    });

    // Expand / Collapse All
    $('.js-expand-collapse').click(function(){
        if($(this).hasClass('js-collapse')) {
            // Collapse
            collapseAllTreeNodes();
        } else {
            // Expand
            expandAllTreeNodes();
        }
    });

    // Cascade and bubble checks
    $("body").delegate('.js-treetable input:checkbox', 'change', function() {
        var val = $(this).prop('checked');
        cascadeChecks($(this), val);
        bubbleChecks($(this), val);
    });

    // Toggle Expand/Collapse
    $("body").delegate('.js-toggle-expand', 'click', function(){
        if($(this).parent().hasClass('children-are-invisible')) {
            $(this).html("&#x25BC;");
            $(this).parent().removeClass('children-are-invisible');
        } else {
            $(this).html("&#x25B6;");
            $(this).parent().addClass('children-are-invisible');
        }
    });
}

// This function is just for demo purposes - remove in production
function getRandomNumber(min, max) {
    return Math.floor(Math.random() * max) + min;
}

/**
 * This function resets all the checks in tree
 * @return void
 */
function resetTreeChecks() {
    $(".js-treetable input:checkbox").prop('checked', false);
}

/**
 * This function expands all tree nodes, sets collapse/expand text to collapse, and sets all > to v
 * @return void
 */
function expandAllTreeNodes() {
    $(".js-expand-collapse").text("Collapse All");
    $(".js-tree-node").removeClass("children-are-invisible");
    $(".js-toggle-expand").html("&#x25BC;");
    $(".js-expand-collapse").removeClass('js-collapse').removeClass('js-expand').addClass('js-collapse');
}

/**
 * This functoin collapses all tree nodes, sets collapse/expand text to expand, and sets all v to >
 * @return void
 */
function collapseAllTreeNodes() {
    $(".js-expand-collapse").text("Expand All");
    $(".js-tree-node").addClass("children-are-invisible");
    $(".js-toggle-expand").html("&#x25B6;");
    $(".js-expand-collapse").removeClass('js-collapse').removeClass('js-expand').addClass('js-expand');
}

/**
 * This function collapses all nodes, resets all checks, and then expands and checks a single node
 * @param  {[type]} id [description]
 * @return {[type]}    [description]
 */
function expandAndCheckTreeNode(id) {
    collapseAllTreeNodes();
    resetTreeChecks();  // Turn this off to expand+reveal multiple items at once, e.g. revealing a previously set list of items
    $(".tree-node#"+id).parents(".tree-node").removeClass('children-are-invisible').children(".toggle-expand").html("&#x25BC;");
    cascadeChecks($(".tree-node#"+id).children('input:checkbox'), true);
    bubbleChecks($(".tree-node#"+id).children('input:checkbox'), true);
}

/**
 * This function cascades a check down to children
 * @param  {jQuery object} checkbox The jQuery tree-node object
 * @param  {boolean}       val      The value of the check
 * @return void
 */
function cascadeChecks(checkbox, val) {
    checkbox.parent().find("input:checkbox").each(function() {
        $(this).prop('checked', val);
    });
}

/**
 * This function bubbles a check up through parents (e.g. if all parent children are checked, check the parent)
 * @param  {jQuery object} checkbox The jQuery tree-node object
 * @param  {boolean}       val      The value of the check
 * @return void
 */
function bubbleChecks(checkbox, val) {
    checkbox.parents('.tree-branch').each(function(){
        if($(this).children('.tree-node').children('.tree-branch').length > 0) {
            var allChecked = $(this).children('.tree-node').children('.tree-branch').find("input:checkbox:not(:checked)").length == 0;
            $(this).children('.tree-node').children('input:checkbox').prop('checked', allChecked);
        }
    });
}