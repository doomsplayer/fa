(function() {
    var app = angular.module('badmintonhome', ['ngResource', 'ngRoute']);
    // URL route
    app.config(['$routeProvider',
        function($routeProvider) {
            $routeProvider.when('/', {
                templateUrl: 'static/tpl/index.html',
                controller: 'IndexCtrl'
            }).when('/about-us', {
                templateUrl: 'static/tpl/about-us.html',
                controller: 'defaultCtrl'
            }).when('/portfolio', {
                templateUrl: "static/tpl/portfolio.html",
                controller: 'PortfolioCtrl'
            }).when('/videoLib',{
                templateUrl: "static/tpl/video-lib.html",
                controller:'VideoLibCtrl'
            })
        }
    ])


    app.controller('PortfolioCtrl', ['$scope',
        function($scope) {
            $rootScope.is_index = false;
        }
    ]).controller('IndexCtrl', ['$rootScope',
        function($rootScope) {
            $rootScope.is_index = true;
        }
    ]).controller('defaultCtrl', ['$rootScope',
        function($rootScope) {
            $rootScope.is_index = false;
        }
    ]).controller('VideoLibCtrl', ['$rootScope', function($rootScope){
        $rootScope.is_index = false;
        
    }]).directive('badmintonLearn', function() {
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
    app.directive('mainBar', function() {
        return {
            scope: {},
            controller: ['$scope', '$http',
                function($scope, $http) {
                    $http.get('/api/common/mainBar').success(function(data, status, headers, config) {
                        $scope.titles = data;
                        setTimeout(function() {
                            $('#idmenu').mnmenu();
                            $('#idmenu').addClass('style_menu');
                        }, 2)

                    })
                }
            ],
            restrict: 'E',
            templateUrl: 'static/tpl/mainBar.html',
            replace: true,
        };
    });

})();