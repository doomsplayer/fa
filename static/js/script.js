(function() {
    var app = angular.module('badmintonhome',['ngResource']);
    app.directive('badmintonLearn', function() {
        return {
            restrict: 'E',
            templateUrl: 'static/tpl/badminton-learn.html'

        }
    });
    // app.directive('badmintonSale', function() {
    //     return {
    //         restrict: 'E',
    //         templateUrl: 'static/tpl/badminton-sale.html'

    //     }
    // });
    app.directive('badmintonVideo', function() {
        return {
            restrict: 'E',
            templateUrl: 'static/tpl/badminton-video.html'

        }
    });
    app.directive('mainBar', function(){
        return {
            scope: {},
            controller: ['$scope','$http',function($scope,$http){
                $http.get('/v1/api/mainBar').success(function(data,status,headers,config){
                    $scope.titles = data;
                    $('#idmenu').mnmenu();
                    $('#idmenu').addClass('style_menu');
                })
            }],
            restrict: 'E',
            templateUrl: 'static/tpl/mainBar.html',
            replace: true,
        };
    });
    app.controller('PortfolioItem', ['$scope', function($scope){
        $scope.items = []
    }])
})();