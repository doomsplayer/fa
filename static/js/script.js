is_index = true;
(function() {
    var app = angular.module('badmintonhome', ['ngResource', 'ngRoute','ui.bootstrap']);
    app.run(['$rootScope','$location',function($rootScope,$location){
        // Fix for nav bar
        $rootScope.index_css = function(){
            if ($location.path() == '/'){
                is_index = true;
                return '';
            }else{
                is_index = false;
                return 'unindex';
            }
        }
    }])
    // URL route
    app.config(['$routeProvider',
        function($routeProvider) {
            $routeProvider.when('/', {
                templateUrl: 'static/tpl/index.html',
                controller: 'defaultCtrl'
            }).when('/about-us', {
                templateUrl: 'static/tpl/about-us.html',
                controller: 'defaultCtrl'
            }).when('/portfolio', {
                templateUrl: "static/tpl/portfolio.html",
                controller: 'PortfolioCtrl'
            }).when('/videoLib',{
                templateUrl: "static/tpl/video-lib.html",
                controller:'VideoLibCtrl'
            }).when('/videoAlbum/:albumId/',{
                templateUrl: "static/tpl/video-album.html",
                controller:'VideoAlbumCtrl'
            })
        }
    ])


    app.controller('$rootScope', ['$', function(){
        
    }]).controller('PortfolioCtrl', ['$scope',
        function($scope) {
        }
    ]).controller('defaultCtrl', ['$rootScope',
        function($rootScope) { }
    ]).controller('VideoLibCtrl', ['$rootScope','$scope', function($rootScope,$scope){
        // FIXME 视频库的API
        $scope.slides = []
        var sample = {pic:'../static/img/video-carousel-0.jpg',title:'赠粉丝——李龙大郑在成搭档这七年',content:'韩国羽毛球名将李龙大/郑在成在北京时间2012年8月5日进行的伦敦奥运会羽毛球男双3、4名争夺战中战胜马来西亚组合获得一枚铜牌。伦敦奥运也是郑在成代表韩国参加的最后的比赛，自此他将正式退役。这枚男双铜牌也将是他运动员生涯的最后一枚奖牌。自此，这对创造无数传奇的的搭档也将正式解散。图为李龙大与郑在成获得铜牌后紧紧相拥。'}
        $scope.slides.push(sample)
        $scope.slides.push(sample)
        $scope.slides.push(sample)
        
        $scope.videos = []
        var temp = []
        var video = {preview:'../static/img/img_test4.jpg',url:'',title:'假动作的魅力：盖得假动作集锦',desc:'描述描述描述描述描述描述描述描述描述描述描述描述描述描述描述描述描述描述描述描述描述描述描述描述描述描述描述描述描述描述描述描述描述描述描述描述描述描述'}
        
        for(var i=1;i<7;i++){
            temp.push(video)
        }
        $scope.videos.push({name:'国际大赛',videos:temp})
        $scope.videos.push({name:'经典对战专辑',videos:temp})
        $scope.videos.push({name:'玩转羽球',videos:temp}) // album.url 

        
    }]).controller('VideoAlbumCtrl', ['$scope','$rootScope','$routeParams', function($scope,$rootScope,$routeParams){
        var id = $routeParams.albumId

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