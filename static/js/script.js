is_index = true;
(function() {
    var app = angular.module('badmintonhome', ['ngResource', 'ngRoute','ngSanitize','ui.bootstrap']);
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
            }).when('/worldChampion',{
                templateUrl: "static/tpl/video-worldchampion.html",  // Seems this page is a part of album ?
                controller: 'WorldChampionCtrl'
            }).when('/learnBadminton',{
                templateUrl: 'static/tpl/learn-to-play-badminton.html',
                controller: 'LearnBadmintonCtrl'
            }).when('/article/:articleId',{
                templateUrl: 'static/tpl/article.html',
                controller: 'ArticleCtrl'
            }).when('/item/:itemId',{
                templateUrl: 'static/tpl/badminton-item.html',
                controller: 'ItemCtrl'
            }).when('/portfolio',{
                templateUrl: 'static/tpl/badminton-sale.html',
                controller: 'PortfolioCtrl'
            }).when('/video/:videoType/:videoId',{
                templateUrl: 'static/tpl/video-play.html',
                controller: 'VideoCtrl'
            }).when('/test',{
                templateUrl: 'static/tpl/test.html'
            })
        }
    ])


    app.controller('PortfolioCtrl', ['$scope',function($scope) {
        var temp = {name:'尤尼克斯YONEX/YY VTZF2 李宗伟最新羽毛球拍SP版TW版',price:'￥1050（正品包邮）',pic:'../static/img/item_list_img.jpg',url:undefined}
        $scope.recommand = [{heading:'最新',items:[temp,temp,temp]},{heading:'即将过期',items:[temp,temp,temp]}]
        
        var temp2 = {}
        temp2.pic = '../static/img/item_main.jpg'
        temp2.type = '球拍'
        temp2.title = '期待已久的弓箭11TH SP版明后天终于要到货啦，超级经典，绝版'
        temp2.price = '￥868(王局体育预订价，包顺丰，限时抢购中)'
        temp2.link = {url:undefined,title:'淘宝',coupon:"ABCDEFG"}
        temp2.performance = performance = ['xxxxxxxxxxxxxxxx','xxxxxxxxxxxxxxxxx','xxxxxxxxxxxxx','xxxxxxxxxxxxxxxxx','xxxxxxxxxxx']
        $scope.items = [temp2,temp2,temp2,temp2,temp2,temp2,temp2]
    }]).controller('defaultCtrl', function() {
    }).controller('VideoLibCtrl', ['$scope', function ($scope) {
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
    }]).controller('VideoAlbumCtrl', ['$scope','$routeParams', function($scope,$routeParams){
        var id = $routeParams.albumId
        // FIXME 按照ID请求专题数据填充
        $scope.name = '2014澳大利亚羽毛球公开赛';
        $scope.desc = '简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介简简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介简简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介简简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介...';
        var video1 = {url:undefined,preview:'../static/img/img_test4.jpg',name:'假动作的魅力：盖得假动作集锦',desc:'描述描述描述描述描述描述描述描述描述描述描述描述描述描述描述描述描述描述描述描述描述描述描述描述描述描述描述描述描述描述描述描述描述描述描述描述描述描述'}
        $scope.related = [video1,video1,video1,video1,video1,video1]
        var video2 = {url:undefined,preview:'../static/img/img_test5.jpg',name:'2013澳大利亚羽毛球公开赛',desc:'简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介简...'}
        $scope.recommand = [video2,video2,video2,video2]
    }]).controller('WorldChampionCtrl', ['$scope', function($scope){
        $scope.title = '国际大赛专辑'
        $scope.desc = '按赛事分类、按时间收录近期在各个国家举办的羽毛球国际赛事精彩视频专辑'
        var item = {pic:'../static/img/img_test3.jpg',url:undefined,desc:'苏迪曼杯，又称世界羽毛球混合团体锦标赛，是羽毛球三大世界性团体赛之一。1989年首届苏迪曼杯在雅加达举办，至今近30年，其中涌现的经典对决不计其数，小编在此为您准备了20场经典巅峰对决，带您一起走进苏迪曼杯这些年。',title:'苏迪曼杯这些年——盘点苏杯经典大战20场'}
        $scope.tabs = []
        $scope.tabs.push({name:'最新',items:[item,item,item,item]})
        $scope.tabs.push({name:'热门',items:[item,item,item]})

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
        var slide = {pic:'../static/img/video-carousel-1.jpg',title:'你真的会缠手胶么？教你缠手胶的小窍门！',desc:'简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介简...'}
        $scope.slides = [slide,slide,slide]
        
        $scope.knowledge = []

        $http.get('/api/common/hottutorial',{params: {type: '大话羽球', from: 0, num: 4}})
             .success(function(response, status, headers, config) {
                if (response.ok) {
                    var hts = response.tutorials;
                    var tmp = [];
                    for (var i in hts) {
                        var ht = hts[i];
                        var picurl = '';
                        $http.get('/api/common/upload',{params: {id: ht.Picid}})
                        .success(function(response, status, headers, config) {
                            picurl = response.filepath;
                            tmp.push(
                            {
                                pic: picurl, 
                                title: ht.Title, 
                                url: '#/article/' + ht.Id, 
                                desc: '描述',
                            }
                            );
                        })
                    }
                    $scope.knowledge.push({name:'大话羽球', link:undefined, passages:tmp});
                }
                
            }
        )
        
        $http.get('/api/common/hottutorial',{params: {type: '羽球战术', from: 0, num: 4}})
             .success(function(response, status, headers, config) {
                if (response.ok) {
                    var hts = response.tutorials;
                    var tmp = [];
                    for (var i in hts) {
                        var ht = hts[i];
                        var picurl = '';
                        $http.get('/api/common/upload',{params: {id: ht.Picid}})
                        .success(function(response, status, headers, config) {
                            picurl = response.filepath;
                            tmp.push(
                            {
                                pic: picurl, 
                                title: ht.Title, 
                                url: '#/article/' + ht.Id, 
                                desc: '描述',
                            }
                            );
                        })
                    }
                    $scope.knowledge.push({name:'羽球战术', link:undefined, passages:tmp});
                }
                
            }
        )
        
        $http.get('/api/common/hottutorial',{params: {type: '羽球技术', from: 0, num: 4}})
             .success(function(response, status, headers, config) {
                if (response.ok) {
                    var hts = response.tutorials;
                    var tmp = [];
                    for (var i in hts) {
                        var ht = hts[i];
                        var picurl = '';
                        $http.get('/api/common/upload',{params: {id: ht.Picid}})
                        .success(function(response, status, headers, config) {
                            picurl = response.filepath;
                            tmp.push(
                            {
                                pic: picurl, 
                                title: ht.Title, 
                                url: '#/article/' + ht.Id, 
                                desc: '描述',
                            }
                            );
                        })
                    }
                    $scope.knowledge.push({name:'羽球技术', link:undefined, passages:tmp});
                }
                
            }
        )

        $http.get('/api/common/hottutorial',{params: {type: '羽球知识', from: 0, num: 4}})
             .success(function(response, status, headers, config) {
                if (response.ok) {
                    var hts = response.tutorials;
                    var tmp = [];
                    for (var i in hts) {
                        var ht = hts[i];
                        var picurl = '';
                        $http.get('/api/common/upload',{params: {id: ht.Picid}})
                        .success(function(response, status, headers, config) {
                            picurl = response.filepath;
                            tmp.push(
                            {
                                pic: picurl, 
                                title: ht.Title, 
                                url: '#/article/' + ht.Id, 
                                desc: '描述',
                            }
                            );
                        })
                    }
                    $scope.knowledge.push({name:'羽球知识', link:undefined, passages:tmp});
                }
                
            }
        )


        $scope.part2 = []
        $http.get('/api/common/hottutorial',{params: {type: '伤病防护', from: 0, num: 4}})
             .success(function(response, status, headers, config) {
                if (response.ok) {
                    var hts = response.tutorials;
                    var tmp = [];
                    for (var i in hts) {
                        var ht = hts[i];
                        var picurl = '';
                        $http.get('/api/common/upload',{params: {id: ht.Picid}})
                        .success(function(response, status, headers, config) {
                            picurl = response.filepath;
                            tmp.push(
                            {
                                pic: picurl, 
                                title: ht.Title, 
                                url: '#/article/' + ht.Id, 
                                desc: '描述',
                            }
                            );
                        })
                    }
                    $scope.knowledge.push({name:'伤病防护', link:undefined, passages:tmp});
                }
                
            }
        )
        $http.get('/api/common/hottutorial',{params: {type: '常见问题', from: 0, num: 4}})
             .success(function(response, status, headers, config) {
                if (response.ok) {
                    var hts = response.tutorials;
                    var tmp = [];
                    for (var i in hts) {
                        var ht = hts[i];
                        var picurl = '';
                        $http.get('/api/common/upload',{params: {id: ht.Picid}})
                        .success(function(response, status, headers, config) {
                            picurl = response.filepath;
                            tmp.push(
                            {
                                pic: picurl, 
                                title: ht.Title, 
                                url: '#/article/' + ht.Id, 
                                desc: '描述',
                            }
                            );
                        })
                    }
                    $scope.knowledge.push({name:'常见问题', link:undefined, passages:tmp});
                }
                
            }
        )

    }]).controller('ItemCtrl', ['$http', '$scope','$routeParams', function($http, $scope,$routeParams){
        $http.get('/api/common/promotion/' + $routeParams.itemId).success(function(response) {
            if (response.ok) {
                var p = response.promotion;
                $http.get('/api/common/upload', {params: {id: p.Picid}}).success(function(response) {
                    if (response.ok) {
                        $scope.type = p.Type
                        $scope.title = p.Title
                        $scope.title2 = p.Title2;
                        $scope.pic = response.filePath;
                        $scope.discription = p.Description;
                        $scope.link = {url: p.url, title: p.Description2}
                        var pic = {thumb:'../static/img/item_thumb.jpg',pic:'../static/img/item_main_pic.jpg'}
                        $scope.pics = [pic,pic,pic,pic]
                        $scope.desc = '文字描述文字描述文字描述文字描述文字描述文字描述文字描述文字描述文字描述文字描述文字描述文字描述文字描述文字描述文字描述文字描述文字描述文字描述文字描述文字描述文字描述文字描述文字描述文字描述文字描述文字描述文字描述文字描述文字描述文字描述文字描述文字描述文字描述文字描述文字描述文字描述文字描述文字描述文字描述文字描述文字描述文字描述文字描述文字描述文字描述文字描述文字描述文字描述文字描述文字描述文字描述文字描述文字描述文字描述文字描述文字描述文字描述文字描述文字描述文字描述文字描述文字描述文字描述文字描述文字描述文字描述文字描述文字描述文字描述文字描述文字描述文字描述文字描述文字描述文字描述文字描述文字描述文字描述文字描述文字描述文字描述'
                        var temp = {name:'尤尼克斯YONEX/YY VTZF2 李宗伟最新羽毛球拍SP版TW版',price:'￥1050（正品包邮）',pic:'../static/img/item_list_img.jpg',url:undefined}
                        $scope.recommand = [temp,temp,temp]


                        $scope.current_pic = $scope.pics[0].pic
                        $scope.selected = 0
                        $scope.switch_pic = function(index){
                            $scope.current_pic = $scope.pics[index].pic
                            $scope.selected = index
                            console.log($scope.current_pic)
                        }
                    }
                })
                
                
            }
        })
    }]).controller('VideoCtrl', ['$scope','$resource','$routeParams', function($scope,$resource,$routeParams){
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
            controller:['$scope','$http',function($scope,$http){
                $scope.Learn = []

                var item = {title:'澳洲赛：林丹逆转夺冠女双演横扫...',url:undefined,pic:'static/img/img_test1.jpg'}
                $scope.Learn.push({name:'大话羽球',pic_news:[item,item,item,item,item,item,item,item],text_news:[item,item,item,item,item,item,item,item,item]})
                $scope.Learn.push({name:'羽球知识',pic_news:[item,item,item,item,item,item,item,item],text_news:[item,item,item,item,item,item,item,item,item]})
                $scope.Learn.push({name:'羽球技术',pic_news:[item,item,item,item,item,item,item,item],text_news:[item,item,item,item,item,item,item,item,item]})
                $scope.Learn.push({name:'羽球战术',pic_news:[item,item,item,item,item,item,item,item],text_news:[item,item,item,item,item,item,item,item,item]})
                $scope.Learn.push({name:'伤病防护',pic_news:[item,item,item,item,item,item,item,item],text_news:[item,item,item,item,item,item,item,item,item]})
            }]
        }
    });
    app.directive('badmintonVideo', function() {
        return {
            scope:{},
            restrict: 'E',
            templateUrl: 'static/tpl/badminton-video.html',
            controller: ['$scope','$http',function($scope,$http){
                $scope.videos = []

                var video  = {preview:'../static/img/img_test3.jpg',url:undefined,title:'苏迪曼杯这些年——盘点苏杯经典大战20场',desc:'苏迪曼杯，又称世界羽毛球混合团体锦标赛，是羽毛球三大世界性团体赛之一。1989年首届苏迪曼杯在雅加达举办，至今近30年，其中涌现的经典对决不计其数，小编在此为您准备了20场经典巅峰对决，带您一起走进苏迪曼杯这些年。'}
                $scope.videos.push({name:'国际大赛专辑',videos:[video,video,video]})
                $scope.videos.push({name:'经典对战专辑',videos:[video,video,video]})
                $scope.videos.push({name:'玩转羽球',videos:[video,video,video]})
            }]
        }
    }).directive('shareTo', function(){
        return {
            restrict: 'EA',
            template: '<span class="pull-right"><img src="../static/img/share.jpg" />分享到<a href="#"><img src="../static/img/weibo.jpg" /></a><a href="#"><img src="../static/img/qzone.jpg" /></a><a href="#"><img src="../static/img/weixin.jpg" /></a></span>'
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
})();