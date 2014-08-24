(function(){
	app = angular.module('badmintonhome');
	app.directive('portfolio', function(){
		return {
			controller: ['$http', '$scope','$element','$attrs',function($http, $scope, $element,$attrs) {
				$scope.items = [];
				$scope.category = {};

				$scope.getColumnNumber = function() { 
					var winWidth = $(window).width()
					if (winWidth > 1200) {
						columnNumber = 5;
					} else if (winWidth > 950) {
						columnNumber = 4;
					} else if (winWidth > 600) {
						columnNumber = 3;
					} else if (winWidth > 400) {
						columnNumber = 2;
					} else {
						columnNumber = 1;
					}
					return columnNumber;
				}
				function Resize(){
					var width = $(window).width();
					var itemWidth = Math.floor(width / $scope.getColumnNumber());
					$('#isotope').find('.element-item').each(function(){
						$(this).css({width:itemWidth + 'px'});
					});
				}
				$(window).on('resize',Resize);
				$scope.$on('$destroy',function(){
					$(window).off('resize');
				})
				$http.get('/api/common/promotiontypes').success(
					function(response) {
						if (response.ok) {
							var p = response.promotiontypes;
							for (var i in p) {
								$scope.category[p[i].Name] = p[i].Name;
								$http.get('/api/common/promotion', {params: {type: p[i].Name, num: 16}}).success(function(response){
									if (response.ok) {
										var items = response.promotions;
										$scope.items = $scope.items.concat(items)
									}
								})
							}
							setTimeout(function(){
								// 绑定filter
								$(window).resize()
								var elm = angular.element('.filters');
								elm.find('*[isotope-filter-by]').bind('click', function (i, e) {
									angular.element('.filters .active').removeClass('active');
									angular.element(this).addClass('active');
									var filter = angular.element(this).attr('isotope-filter-by');
									$scope.$broadcast('filter-elements', { filter: filter });
								});
								var option = {
									itemSelector:'.element-item',
									layoutMode:'fitRows'
								}
								// 绑定Item与事件
								var elm2 = angular.element('#isotope')
								elm2.isotope(option);
								$scope.$watch('items',function(){
									setTimeout(function(){
										elm2.isotope('reloadItems').isotope(option)
									})
								},true);
								$scope.$on('filter-elements',function(data,filter){
									option.filter = filter.filter;
									elm2.isotope(option)
								});
							},500)
						}
				});
			}],
			templateUrl: 'static/tpl/portfolio_tpl.html',
			replace: true,
			scope:{}
		};
	});
})();