(function(){
	app = angular.module('badmintonhome');
	app.directive('portfolio', function(){
		return {
			controller: ['$http', '$scope','$element','$attrs',function($http, $scope, $element,$attrs) {
				$scope.items = [];
				$scope.category = {};

				$http.get('/api/common/promotiontypes').success(
					function(response) {
						if (response.ok) {
							var p = response.promotiontypes;
							for (var i in p) {
								$scope.category[p[i].Name] = p[i].Name;
								$http.get('/api/common/promotion', {params: {type: p[i].Name, num: 16}}).success(function(response){
									if (response.ok) {
										var items = response.promotions;
										for (var i in items) {
											$http.get("/api/common/upload", {params: {id: items[i].PicId}}).success(function(response){
												if (response.ok){
													$scope.items.push(
													{	
														name:items[i].Title, 
														pic_url: response.filepath, 
														price: items[i].Description2, 
														category: items[i].Type,
														id: items[i].Id
													})
												}
												
											})
											
										}
										
									}
									
								})
								
							}
						}
				});

				setTimeout(function(){
                    // 绑定filter
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
                    },true)
                    $scope.$on('filter-elements',function(data,filter){
                    	option.filter = filter.filter;
                    	elm2.isotope(option)
                    })

				},0)

			}],
			templateUrl: 'static/tpl/portfolio_tpl.html',
			replace: true,
			scope:{}
		};
	});
})();