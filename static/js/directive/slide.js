(function(){
	var app = angular.module('badmintonhome')
	app.directive('slider',function(){
		return {
			scope: {}, // {} = isolate, true = child, false/undefined = no change
			controller: ['$scope','$element',function($scope, $element) {
				$scope.slides = []
				// DELETEME
				$scope.slides.push({pic_url:'static/img/img1.jpg',pic_alt:'image01',title:'欢迎来到羽球之家',subtitle:'副标题'})
				$scope.slides.push({pic_url:'static/img/img0.jpg',pic_alt:'image02',title:'欢迎来到羽球之家',subtitle:'副标题'})
				$scope.slides.push({pic_url:'static/img/img1.jpg',pic_alt:'image03',title:'欢迎来到羽球之家',subtitle:'副标题'})
				$scope.slides.push({pic_url:'static/img/img0.jpg',pic_alt:'image04',title:'欢迎来到羽球之家',subtitle:'副标题'})
				//  FIXME 从远端Fetch
				setTimeout(function(){
					angular.element($element[0].firstElementChild).eislideshow({
						animation			: 'center',
						autoplay			: true,
					 	slideshow_interval	: 5000,
						titlesFactor		: 0
					});
				},0)

			}],
			restrict: 'A', // E = Element, A = Attribute, C = Class, M = Comment
			templateUrl: 'static/tpl/slide.html',
			replace: true
		};
	});
	app.directive('backstretch', function(){
		return {
			scope: {},
			controller: ['$scope','$element',function($scope, $element) {
				$scope.pic_url = 'static/img/head_bg.jpg' // FIXME 是否与服务器同步 ？
				setTimeout(function(){
					angular.element($element[0]).backstretch($scope.pic_url)
				},0)
			}],
			restrict: 'AE',
			template: '<div class="background backstretch" data-background-image="{{pic_url}}"></div>',
			replace: true
		};
	});
})();