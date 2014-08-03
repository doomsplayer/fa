(function() {
    var app = angular.module('badmintonhome', []);

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
            // scope: {}, // {} = isolate, true = child, false/undefined = no change
            controller: ['$scope',function($scope){
                // FIXME 使用$resource 替代
                $scope.titles = [{title:'促销信息',url:'#'},{title:'学打羽毛球',url:'#',subtitle:[{'url': '#', 'title': '大话羽球'}, {'url': '#', 'title': '羽球知识'}, {'url': '#', 'title': '羽球技术'}, {'url': '#', 'title': '羽球战术'}, {'url': '#', 'title': '伤病防护'}]},{'title':'视频库','url':'#','subtitle':[{'url': '#', 'title': '国际大赛专辑'}, {'url': '#', 'title': '经典对战专辑'}, {'url': '#', 'title': '玩转羽球'}]},{'url':'#','title':'联系我们'}]
                setTimeout(function(){
                    $('#idmenu').mnmenu();
                    $('#idmenu').addClass('style_menu');
                },0)
                // workaround ...
                
            }],
            // require: 'ngModel', // Array = multiple requires, ? = optional, ^ = check parent elements
            restrict: 'E', // E = Element, A = Attribute, C = Class, M = Comment
            templateUrl: 'static/tpl/mainBar.html',
            replace: true,
            // transclude: true,
            // compile: function(tElement, tAttrs, function transclude(function(scope, cloneLinkingFn){ return function linking(scope, elm, attrs){}})),
            link: function($scope, iElm, iAttrs, controller) {
                
            }
        };
    });
})();