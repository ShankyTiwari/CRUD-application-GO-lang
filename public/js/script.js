'use strict';

const app = angular.module('app',[]);

app.service('appService', function ($http) {	
	return {
		getUsers : function() {
			return new Promise( function (resolve, reject) {
				$http({
					method: 'GET',
					url: '/user',
					headers: {
						'Content-Type': undefined
					},
				})
				.then( function(response){
					resolve(response.data);
				})
				.catch(function (error) {
					reject([]);
				})
			});
		},
		addUser: function (requestParamter, userId) {
			return new Promise(function (resolve, reject) {
				if( userId !== null ) {
					requestParamter.id = userId;
				}
				$http({
					method: userId === null ? 'POST' : 'PUT' ,
					url: '/user',
					headers: {
						'Content-Type': undefined
					},
					data: requestParamter
				})
				.then(function (response) {
					resolve(response.data);
				})
				.catch(function (error) {
					reject([]);
				})
			});
		},
		deleteUser: function (userId) {
			return new Promise(function (resolve, reject) {
				$http({
					method: 'DELETE',
					url: '/user/' + userId,
					headers: {
						'Content-Type': undefined
					}
				})
				.then(function (response) {
					resolve(response.data);
				})
				.catch(function (error) {
					reject([]);
				})
			});
		}
	}
});

app.controller('app', function($scope,appService) {	

	$scope.selectedUserId = null;

	appService.getUsers()
		.then(function (users) {
			$scope.userList = users;
			setTimeout(() => {
				$scope.$apply();
			});
		})
		.catch(function (error) {
			console.log(error);
		});
	
	$scope.addUser = () => {
		if ($scope.name === '' || typeof $scope.name === 'undefined') {
			alert(`Enter the Name.`);
		} else if ($scope.lname === '' || typeof $scope.lname === 'undefined') {
			alert(`Enter the Last Name.`);
		} else if ($scope.country === '' || typeof $scope.country === 'undefined') {
			alert(`Enter the Country.`);
		} else {
			appService.addUser(
				{
					name: $scope.name,
					lname: $scope.lname,
					country: $scope.country
				},
				$scope.selectedUserId
			)
			.then(function (users) {
				$scope.userList = users;
				$scope.$apply();
			})
			.catch(function (error) {
				console.log(error);
			});
		}
	}; 
	
	$scope.deleteUser = function (userId) {
		if (userId === '' || userId === undefined) {
			alert(`User ID is empty`);
		} else {
			appService.deleteUser(userId)
			.then(function (users) {
				$scope.userList = users;
				$scope.reset();
				$scope.$apply();
			})
			.catch(function (error) {
				console.log(error);
			});
		}
	};

	$scope.selectUser = function (userDetail) {
		$scope.selectedUserId = userDetail.ID;
		$scope.name = userDetail.Name;
		$scope.lname = userDetail.Lname;
		$scope.country = userDetail.Country;
	}

	$scope.reset = function () {
		$scope.selectedUserId = null;
		$scope.name = null;
		$scope.lname = null;
		$scope.country = null;
	}
	
});