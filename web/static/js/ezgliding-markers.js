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
// Helper functions for marker creation.
//
// Relies on the existence of two globals for marker handling:
// 	mgrAirfield and mgrWaypoint
//
// Check ezgliding-map.js for their declaration.
//

// addMarkers creates adds the feature's markers to the markermanager (mgr).
// feature can represent any of an airfield or waypoint.
function addMarkers(feature) {
	var goType = feature.properties.Go;
	switch(goType) {
	case 'Airfield':
		var flags = feature.properties.Flags;

		if(flags & Outlanding) {
			addOutlandingMarkers(feature);
		} else {
			addAirfieldMarkers(feature);
		}
		break;
	case 'Waypoint':
		addWaypointMarkers(feature);
		break;
	}
}

// addAirfieldMarkers adds the airport feature's markers to the markermanager.
// it should handle only feature with no 'Outlanding' flag set.
function addAirfieldMarkers(feature) {
	var flags = feature.properties.Flags;
	var color = (flags & Asphalt || flags & Concrete) ? '#336699' : '#550055';

	var icon = { 
		path: "M-10,0 a10,10 0 1,0 20,0 a10,10 0 1,0 -20,0 M-8,-1 L8,-1 L8,1 L-8,1 L-8,-1",
		strokeWeight: 0, fillColor: color, fillOpacity: 0.7 };

	var runway = feature.properties.Runway;
	if(runway.length==4) {
		icon.rotation = (parseInt(runway.substring(0, 2))*10) - 90;
	}
	icon.scale = 1;

     	var mkA = new MarkerWithLabel({
       		position: feature2LatLng(feature),
      		icon: icon,
       		labelContent: getAirfieldText(feature),
       		labelAnchor: new google.maps.Point(0, -12),
       		labelClass: "airportlabel",
       		labelStyle: {opacity: 0.75, color: color}
     	});

        mgrAirfield.addMarker(mkA, 10);
		
	return [mkA];
}

// getAirfieldText returns the label for the given Airfield feature.
function getAirfieldText(feature) {
	var text = feature.properties.Name;
       	
	if (feature.properties.ICAO != '') {
		text += ' (' + feature.properties.ICAO + ')';
	}
	text += '<br/>';

	if (feature.properties.Frequency != '0') {
		text += feature.properties.Frequency + ' '
	} else {
		text += ' ';
	}
	text += feature.properties.Elevation + 'm';

	return text;
}

// addOutlandingMarkers adds the outlanding feature's markers to the markermanager.
// it should handle only feature with 'Outlanding' flag set.
function addOutlandingMarkers(feature) {
	var text = feature.properties.Name;
	if (feature.properties.Catalog != '') {
		text += ' ' + feature.properties.Catalog;
	}
	text += '<br/>' + feature.properties.Elevation + 'm';

	var icon = { 
		path: "M-5,0 a5,5 0 1,0 10,0 a5,5 0 1,0 -10,0",
		strokeWeight: 1, strokeColor: '#1A334C', strokeOpacity: 0.7 };
	icon.scale = 1;

     	var mkA = new MarkerWithLabel({
       		position: feature2LatLng(feature),
      		icon: icon,
       		labelContent: text,
       		labelAnchor: new google.maps.Point(0, -8),
       		labelClass: "outlandinglabel",
       		labelStyle: {opacity: 0.75, color: '#1A334C'}
     	});

        mgrAirfield.addMarker(mkA, 10);

	return [mkA];
}

// addWaypointMarkers adds the waypoint feature's markers to the markermanager.
function addWaypointMarkers(feature) {
	var text = feature.properties.Description;

     	var mkW = new google.maps.Marker({
       		position: feature2LatLng(feature),
      		icon: { path: google.maps.SymbolPath.CIRCLE, strokeWeight: 2,
			scale: 2, strokeColor: 'blue', strokeOpacity: 0.7, fillWeight: 0 },
		title: text
     	});

        mgrWaypoint.addMarker(mkW, 10);

	return [mkW];
}

