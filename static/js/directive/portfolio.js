(function(){
	app = angular.module('badmintonhome');
	app.directive('portfolio', function(){
		return {
			controller: ['$scope','$element','$attrs',function($scope, $element,$attrs) {
				$scope.items = []
				// FIXME 在线Fetch
				$scope.items.push({name:'产品名称',pic_url:'static/img/img_test2.jpg',price:'价格xx元',category:'shoe'})
				$scope.items.push({name:'产品名称',pic_url:'static/img/img_test2.jpg',price:'价格xx元',category:'shoe'})
				$scope.items.push({name:'产品名称',pic_url:'static/img/img_test2.jpg',price:'价格xx元',category:'cloth'})
				$scope.items.push({name:'产品名称',pic_url:'static/img/img_test2.jpg',price:'价格xx元',category:'shoe'})
				$scope.items.push({name:'产品名称',pic_url:'static/img/img_test2.jpg',price:'价格xx元',category:'shoe'})
				$scope.items.push({name:'产品名称',pic_url:'static/img/img_test2.jpg',price:'价格xx元',category:'shoe'})
				$scope.items.push({name:'产品名称',pic_url:'static/img/img_test2.jpg',price:'价格xx元',category:'shoe'})
				$scope.items.push({name:'产品名称',pic_url:'static/img/img_test2.jpg',price:'价格xx元',category:'shoe'})
				$scope.items.push({name:'产品名称',pic_url:'static/img/img_test2.jpg',price:'价格xx元',category:'cloth'})
				$scope.items.push({name:'产品名称',pic_url:'static/img/img_test2.jpg',price:'价格xx元',category:'cloth'})
				$scope.category = {'shoe':'鞋子','cloth':'衣服'}
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
			scope:{},
			link: function($scope, iElm, iAttrs, controller) {
				
			}
		};
	});
})();