(function() {
    var app = angular.module('badmintonhome', []);

    app.directive('badmintonLearn', function() {
        return {
            restrict: 'E',
            templateUrl: '../static/tpl/badminton-learn.html'

        }
    });
    // app.directive('badmintonSale', function() {
    //     return {
    //         restrict: 'E',
    //         templateUrl: '../static/tpl/badminton-sale.html'

    //     }
    // });
    app.directive('badmintonVideo', function() {
        return {
            restrict: 'E',
            templateUrl: '../static/tpl/badminton-video.html'

        }
    });
})();