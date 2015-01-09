// Copyright 2014 The ezgliding Authors.
//
// This file is part of ezgliding.
//
// ezgliding is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// ezgliding is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with ezgliding.  If not, see <http://www.gnu.org/licenses/>.
//
// Author: Ricardo Rocha <rocha.porto@gmail.com>
//
// Handles the map display, including data initialization and marker loading.
//

// map is the global map element.
var map;

// mgrAirfield is the marker manager for Airfield data.
// (http://google-maps-utility-library-v3.googlecode.com/svn/tags/markermanager/1.0/docs/reference.html)
var mgrAirfield;

// mgrAirfield is the marker manager for Waypoint data.
// (http://google-maps-utility-library-v3.googlecode.com/svn/tags/markermanager/1.0/docs/reference.html)
var mgrWaypoint;

// initialize creates and enables the map element, loading required data.
function initialize() {

	var featureOpts = [
		{
    			elementType: 'labels',
			stylers: [
				{ visibility: 'off' }
			]
		},
	];

	mapOptions = {
		zoom: 10,
		center: {lat: 46.2, lng: 6.082},
		mapTypeControlOptions: {
			mapTypeIds: [google.maps.MapTypeId.TERRAIN, google.maps.MapTypeId.SATELLITE]
		},
		mapTypeId: google.maps.MapTypeId.TERRAIN
	};
	map = new google.maps.Map(
			document.getElementById('map-canvas'), mapOptions);
	map.setOptions({styles: featureOpts});

	mgrAirfield = new MarkerManager(map);
	mgrWaypoint = new MarkerManager(map);

	var xhr = new XMLHttpRequest();
	xhr.open('GET', '/airfield/?region=CH&region=FR&accept=application/json', true);
	xhr.onload = function() {
  		loadGeoJSON(JSON.parse(this.responseText));
		mgrAirfield.refresh();
	};
	xhr.send();

	var xhr = new XMLHttpRequest();
	xhr.open('GET', '/waypoint/?region=CH&accept=application/json', true);
	xhr.onload = function() {
  		loadGeoJSON(JSON.parse(this.responseText));
		mgrWaypoint.refresh();
	};
	xhr.send();
}
google.maps.event.addDomListener(window, 'load', initialize);

// loadGeoJSON is an alternative to v3's loadGeoJSON, adding markers individually.
// reasoning to have this custom function is to have more flexibility on the type
// of markers being loaded, as well as having multiple markers for a single point.
function loadGeoJSON(json) {
	for (i=0; i<json.features.length; i++) {
		var feature = json.features[i];
		addMarkers(feature);
	}
}

