'use strict';

const app = angular.module('app',[]);

/* 
    Making factory method for socket 
*/
app.service('ajax', function($http,$httpParamSerializer){

	return {
		execute : (data,callback) => {
			$http.post(data.url,(data.params)).then( function(data, status, headers, config){
		    	callback(data);
		    }).catch( function(data, status) {
		        alert("Ajax Connection Error");
		    });
		}
	}
});

app.service('appService', function(ajax){
	
	return {
		getUsers : ( callback ) => {
			ajax.execute({
					url : '/getUsers',
					params : {}
				},( response )=>{
					console.log(response.data)
					callback( response.data);
				}
			);
		},
		insertUser : (data,callback) => {
			ajax.execute( data ,( response )=>{
					callback( response.data);
			});
		}
	}

});

app.controller('app', function($scope,appService){

	appService.getUsers( (users) =>{
		$scope.userList = users;
	});	

	$scope.insertUser = () => {

		if ($scope.name === '' || typeof $scope.name ==='undefined') {
			
			alert( `Enter the Name.` );

		} else if($scope.lname === '' || typeof $scope.lname ==='undefined') {
			
			alert( `Enter the Last Name.` );
			
		}else if($scope.country === '' || typeof $scope.country ==='undefined') {
			
			alert( `Enter the Country.` );
			
		}else{

			const data = {
				url : '/insertUsers',
				params : {
					name : $scope.name,
					lname : $scope.lname,
					country : $scope.country,
				}
			};
			appService.insertUser(data,( users )=>{
				$scope.userList = users;
			});
		}
	};

});