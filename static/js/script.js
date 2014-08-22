is_index = true;
(function() {
    var app = angular.module('badmintonhome', ['ngResource', 'ngRoute','ngSanitize','ui.bootstrap','infinite-scroll']);
    app.run(['$rootScope','$location',function($rootScope,$location){
        // Fix for nav bar
        $rootScope.index_css = function(){
            if ($location.path() == '/'){
                is_index = true;
                $rootScope.is_index = true;
                return '';
            }else{
                is_index = false;
                $rootScope.is_index = false;
                return 'unindex';
            }
        };
        $rootScope.index = function(){
            $location.path('/');
        }

    }])
    
    app.config(function($httpProvider) {
      // Use x-www-form-urlencoded Content-Type
      $httpProvider.defaults.headers.post['Content-Type'] = 'application/x-www-form-urlencoded;charset=utf-8';
     
      /**
       * The workhorse; converts an object to x-www-form-urlencoded serialization.
       * @param {Object} obj
       * @return {String}
       */ 
      var param = function(obj) {
        var query = '', name, value, fullSubName, subName, subValue, innerObj, i;
          
        for(name in obj) {
          value = obj[name];
            
          if(value instanceof Array) {
            for(i=0; i<value.length; ++i) {
              subValue = value[i];
              fullSubName = name + '[' + i + ']';
              innerObj = {};
              innerObj[fullSubName] = subValue;
              query += param(innerObj) + '&';
            }
          } else if(value instanceof Object) {
            for(subName in value) {
              subValue = value[subName];
              fullSubName = name + '[' + subName + ']';
              innerObj = {};
              innerObj[fullSubName] = subValue;
              query += param(innerObj) + '&';
            }
          } else if(value !== undefined && value !== null) {
            query += encodeURIComponent(name) + '=' + encodeURIComponent(value) + '&';
          }
        }
          
        return query.length ? query.substr(0, query.length - 1) : query;
      };
     
      // Override $http service's default transformRequest
      $httpProvider.defaults.transformRequest = [function(data) {
        return angular.isObject(data) && String(data) !== '[object File]' ? param(data) : data;
      }];
    });

    // URL route
    app.config(['$routeProvider',
        function($routeProvider) {
            $routeProvider.when('/', {
                templateUrl: 'static/tpl/index.html',
                controller: 'defaultCtrl'
            }).when('/about-us', {
                templateUrl: 'static/tpl/about-us.html',
                controller: 'defaultCtrl'
            }).when('/videoLib',{
                templateUrl: "static/tpl/video-lib.html",
                controller:'VideoLibCtrl'
            }).when('/videoAlbum/:albumId',{
                templateUrl: "static/tpl/video-album.html",
                controller:'VideoAlbumCtrl'
            }).when('/video/:videoId',{
                templateUrl: 'static/tpl/video-play.html',
                controller: 'VideoCtrl'
            }).when('/learnBadminton',{
                templateUrl: 'static/tpl/article-index.html',
                controller: 'LearnBadmintonCtrl'
            }).when('/articleAlbum/:albumId',{
                templateUrl: 'static/tpl/article-list.html',
                controller: 'ArticleAlbumCtrl'
            }).when('/article/:articleId',{
                templateUrl: 'static/tpl/article.html',
                controller: 'ArticleCtrl'
            }).when('/item/:itemId',{
                templateUrl: 'static/tpl/badminton-item.html',
                controller: 'ItemCtrl'
            }).when('/portfolio',{
                templateUrl: 'static/tpl/portfolio.html',
                controller: 'PortfolioCtrl'
            }).when('/test',{
                templateUrl: 'static/tpl/test.html'
            })
        }
    ])
    
    app.factory('Api', ['$resource', function($resource){
        var api = {};
        api.file = $resource('/api/common/upload',{},{get:{cache:true,method:'GET'}});
        api.video = $resource('/api/common/video',{},{get:{cache:true,method:'GET'}});
        api.tutorialtypes = $resource('/api/common/tutorialtypes', {}, {get:{cache:true,method:'GET'}});
        api.tutorial = $resource('/api/common/tutorial')
        return api;
    }])

    app.controller('PortfolioCtrl', ['$http', '$scope','$log',function($http, $scope, $log) {
        $scope.now = 0;
        $scope.msg = '更多促销信息加载中，请稍候'
        $scope.loading = true;
        $http.get('/api/common/promotion', {params: {num: 20}}).success(function(response) {
            if (response.ok) {
                var items = response.promotions;
                $scope.loading = false;
                $scope.items = items;
                $scope.now += items.length;
                $scope.recommand = [{heading:'最新',items: items},{heading:'即将过期',items: items}]    
            }
        });
        $scope.nextPage = function(){
            if ($scope.loading) return;
            $scope.loading = true;
            $http.get('/api/common/promotion',{params:{num:20,from:$scope.now}}).success(function(response){
                if(response.ok){
                    var items = response.promotions;
                    for (var i in items){
                        $scope.items.push(items[i]);
                    }
                    $scope.now += items.length;
                    $log.log('load ' + items.length + ' promotions');
                    $scope.loading = false;
                }else{
                    if (response.errmsg == 'empty'){
                        $log.log('No more data')
                        $scope.msg = '您已经全部看完啦，休息一下吧！'
                    }
                }
            })
        }
        
    }]).controller('defaultCtrl', function() {
    }).controller('VideoLibCtrl', ['$http', '$scope', function($http, $scope) {
        $http.get('/api/common/videotypes').success(function(resp) {
            $scope.albums = [];
            if (resp.ok) {
                for (var i in resp.videotypes) {
                    var typeName = resp.videotypes[i].Name;

                    $http.get('/api/common/hotvideo', {params: {type: typeName, num: 6}})
                    .success(function(typeName) {
                        return function(resp) {
                            if (resp.ok) {
                                console.log(typeName)
                                $scope.albums.push({name: typeName,videos: resp.videos})
                            }
                        }
                    }(typeName))
                    
                }
            }
        })
        $http.get('/api/common/hotvideo', {params: {num: 3}}).success(function(r) {
            if (r.ok) {
                $scope.slides = r.videos;
            }
        })
        
        

    }]).controller('VideoAlbumCtrl', ['$http', '$scope','$routeParams', '$log', function($http, $scope, $routeParams, $log){
        var id = $routeParams.albumId;
        $scope.now = 0
        $scope.loading = true
        $http.get("/api/common/video", {params: {type: id, num: 30}}).success(function(r) {
            console.log(r)
            $scope.name = id;
            $scope.desc = '';
            if (r.ok) {
                $scope.related = r.videos;    
                $scope.recommand = r.videos;
                $scope.now += r.videos.length;
                $scope.loading = false;
            } 
        })
        $scope.nextPage = function(){
            if ($scope.load) return;
            $scope.loading = true;
            $http.get("/api/common/video", {params: {type: $routeParams.albumId, num: 30, from: $scope.now}}).success(function(r){
                if (r.ok){
                    for(var i in r.videos){
                        $scope.related.push(r.videos[i]);
                        // $scope.recommand.push(r.videos[i]);  //FIXME Should recommands be infinite-Scollable ?
                    }
                    $scope.now += r.videos.length;
                    $log.log('Load ' + r.videos.length + ' videos');
                    $scope.loading = false;
                }else{
                    if (r.errmsg == 'empty'){
                        $log.log('No more data');
                        $scope.msg = '您已经全部看完啦，休息一下吧！';
                    }
                }
            })
        }

    }]).controller('ArticleAlbumCtrl', ['$http', '$scope', '$routeParams', '$log', function($http, $scope, $routeParams, $log){
        var id = $routeParams.albumId;
        $scope.msg = '更多促销信息加载中，请稍后'
        $scope.now = 0
        $scope.loading = true
        $http.get('/api/common/tutorial',{params: {type: id, num: 20}}).success(function(r) {
            if (r.ok) {
                $scope.title = id;
                $scope.desc = '按赛事分类、按时间收录近期在各个国家举办的羽毛球国际赛事精彩视频专辑';
                $scope.tabs = [];
                $scope.tabs.push({name:'最新',items: r.tutorials});
                $scope.tabs.push({name:'热门',items: r.tutorials});
                $scope.now += r.tutorials.length;
                $scope.loading = false;
            }
        });
        $scope.nextPage = function(){
            if ($scope.loading) return;
            $scope.loading = true
            $http.get('/api/common/tutorial',{params: {type:$routeParams.albumId, num:20, from:$scope.now}}).success(function(r){
                if(r.ok){
                    for(var i in r.tutorials){
                        $scope.tabs[0].items.push(r.tutorials[i]);
                        $scope.tabs[1].items.push(r.tutorials[i]);
                    };
                    $scope.now += r.tutorials.length;
                    $scope.loading = false;
                    $log.log('load ' + r.tutorials.length + ' Articles');
                }else{
                    if (r.errmsg == 'empty'){
                        $log.log('No more data')
                        $scope.msg = '您已经全部看完啦，休息一下吧！'
                    }
                }
            })
        }
        

    }]).controller('ArticleCtrl', ['$http', '$routeParams','$scope', function($http, $routeParams,$scope){
        var articleId = $routeParams.articleId
        $http.get('/api/common/tutorial/' + articleId).success(function(response, status, headers, config){
            if (response.ok) {
                var tutorial = response.tutorial;
                $scope.author = tutorial.Author;
                $scope.source = tutorial.Source;
                $scope.date = tutorial.Time;
                $scope.click = tutorial.Click;
                $scope.star = tutorial.Favor;
                $scope.title = tutorial.Title
                $scope.content = tutorial.Content
                $http.get('/api/common/hottutorial?type=' + tutorial.Type).success(function(response, status, headers, config){
                    if (response.ok) {
                        var ts = response.tutorials;
                        $scope.recommand = []
                        for (var i in ts) {
                            $scope.recommand.push({title: ts[i].Title, date: ts[i].Time ,url: "#/article/" + ts[i].Id})
                        }
                        
                    }
                })
            }
        })
    }]).controller('LearnBadmintonCtrl', ['$http', '$scope', function($http, $scope){
        $http.get('/api/common/hottutorial', {params: {num: 3}}).success(function(r) {
            if (r.ok) {
                $scope.slides = r.tutorials;
            }
        });
        $scope.knowledge = []
        $http.get('/api/common/tutorialtypes').success(function(r) {
            if (r.ok) {
                for (var i in r.tutorialtypes) {   
                    $http.get('/api/common/hottutorial',{params: {type: r.tutorialtypes[i].Name, from: 0, num: 4}})
                         .success(
                            function(i) {
                                return function(response, status, headers, config) {
                                    if (response.ok) {
                                        $scope.knowledge.push({name: r.tutorialtypes[i].Name, link: '/#/articleAlbum/' + r.tutorialtypes[i].Name, passages: response.tutorials});
                                    }
                                }
                            }(i)
                        )
                }
            }
        })
    }]).controller('ItemCtrl', ['$http', '$scope','$routeParams', function($http, $scope,$routeParams){
        $http.get('/api/common/promotion/' + $routeParams.itemId).success(function(response) {
            if (response.ok) {
                $scope.promotion = response.promotion;
                // TODO fetch below data
                var pic = {thumb:'../static/img/item_thumb.jpg',pic:'../static/img/item_main_pic.jpg'};
                $scope.pics = [pic,pic,pic,pic];
                var temp = {name:'尤尼克斯YONEX/YY VTZF2 李宗伟最新羽毛球拍SP版TW版',price:'￥1050（正品包邮）',pic:'../static/img/item_list_img.jpg',url:undefined};
                $scope.recommand = [temp,temp,temp];


                $scope.current_pic = $scope.pics[0].pic;
                $scope.selected = 0;
                $scope.switch_pic = function(index){
                    $scope.current_pic = $scope.pics[index].pic;
                    $scope.selected = index;
                    console.log($scope.current_pic);
                }
            }
        })
    }]).controller('VideoCtrl', ['$http', '$scope','$resource','$routeParams', function($http, $scope,$resource,$routeParams){
        var videoId = $routeParams.videoId;
        var video = $resource('api/common/video')
        var ret = video.get({type:$routeParams.videoType,from:$routeParams.videoId})
        ret.$promise.then(function(){
            $scope.video = ret.videos[0]
            $('.youku-video')[0].innerHTML = $scope.video.Content.replace(/(width|height)=\d+/gi,'');
        })
    }]).directive('badmintonLearn', function() {
        return {
            restrict: 'E',
            templateUrl: 'static/tpl/badminton-learn.html',
            scope:{},
            controller:['$scope','Api','$log',function($scope,Api,$log){
                $scope.Learn = [];
                var types = Api.tutorialtypes.get()
                types.$promise.then(function(){
                    if (types.ok){
                        for (var i in types.tutorialtypes){
                            $scope.Learn.push({name:types.tutorialtypes[i].Name,pic_news:undefined,text_news:undefined})
                        }
                        $scope.click(0);
                    }
                })
                $scope.click = function(num){
                    if (!$scope.Learn[num].pic_news){
                        $log.log('load ' + $scope.Learn[num].name)
                        var d = Api.tutorials.get({type:$scope.Learn[num].name,num:8})
                        d.$promise.then(function(){
                            if (d.ok){
                                $scope.Learn[num].pic_news = d.tutorials;
                                $scope.Learn[num].text_news = d.tutorials;
                            }
                        })
                    }
                }

            }]
        }
    });
    app.directive('badmintonVideo', function() {
        return {
            scope:{},
            restrict: 'E',
            templateUrl: 'static/tpl/badminton-video.html',
            controller: ['$scope','$resource','$http','$q','Api',function($scope,$resource,$http,$q,Api){
                $scope.videos = []
                $http.get('/api/common/videotypes').success(function(response){
                    if (response.ok){
                        var promises = []
                        for(var i in response.videotypes){
                            var Name = response.videotypes[i].Name
                            var ret = Api.video.get({num:3,type:Name})
                            ret.$promise.Name = Name
                            promises.push(ret.$promise)
                        }
                        $q.all(promises).then(function(data){
                            for (var i in data){
                                if (data[i].ok){
                                    $scope.videos.push({name:data[i].$promise.Name,videos:data[i].videos})
                                }
                            }
                        })
                    }
                })
            }]
        }
    }).directive('shareTo', function(){
        return {
            // TODO 引入第三方社会化分享工具
            restrict: 'EA',
            template: '<span class="pull-right"><span class="icon icon-share"></span>分享到<a href="#"><span class="icon icon-weibo"></span></a><a href="#"><span class="icon icon-qzone active"></span></a><a href="#"><span class="icon icon-wechat"></span></a></span>'
        };
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
                        }, 0)

                    })
                }
            ],
            restrict: 'E',
            templateUrl: 'static/tpl/mainBar.html',
            replace: true,
        };
    });
    app.directive('pic',function(){
        return {
            scope:{Id:'@picid'},
            template: '<img class="pic" ng-src="{{url}}"></img>',
            replace: true,
            restrict: 'E',
            controller:['$scope','$element','Api',function($scope,$element,Api){
                ret = Api.file.get({id:$scope.Id})
                ret.$promise.then(function(data){
                    $scope.url = data.filepath
                    // $($element).attr('src',data.filepath)
                })
                $scope.$watch('Id',function(newValue,oldValue){
                    if (newValue != oldValue){
                        ret = Api.file.get({id:newValue})
                        ret.$promise.then(function(data){
                            $scope.url = data.filepath
                            // $($element).attr('src',data.filepath)
                        })
                    }
                })                
            }]
        }
    })

})();