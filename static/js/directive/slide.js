(function(){
	var app = angular.module('badmintonhome')
	app.directive('slider',function(){
		return {
			scope: {}, // {} = isolate, true = child, false/undefined = no change
			controller: ['$scope','$element','$resource','$q','Api',function($scope, $element,$resource,$q,Api) {
				var s = $resource('/api/common/carousel')
				$scope.slides = []
				$scope.fetchUrl = function(item){
					var d = $q.defer();
					var result = Api.file.get({id:item['PicId']})
					result.$promise.then(function(){
						if (result.ok){
							item.PicUrl = result.filepath
							d.resolve(item)
							// console.log(item)
						}
					})
					return d.promise
				}
				ret = s.get({num:4}) // 在这里修改出现的个数
				var waitlist = []
				ret.$promise.then(function(){
					if (ret.ok){
						// console.log(ret.carousels)
						for (var i in ret.carousels){
							waitlist.push($scope.fetchUrl(ret.carousels[i]))
						}
						$q.all(waitlist).then(function(data){
							$scope.slides = data
							setTimeout(function(){
								$($element[0].firstElementChild).eislideshow({
									animation			: 'center',
									autoplay			: true,
									slideshow_interval	: 5000,
									titlesFactor		: 0
								});
							},0)
						})
					}
				})
				
			}],
			restrict: 'A',
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