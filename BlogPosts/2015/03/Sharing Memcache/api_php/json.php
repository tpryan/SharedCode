<!--
	Copyright 2015, Google, Inc. 
 Licensed under the Apache License, Version 2.0 (the "License"); 
 you may not use this file except in compliance with the License. 
 You may obtain a copy of the License at 
  
    http://www.apache.org/licenses/LICENSE-2.0 
  
 Unless required by applicable law or agreed to in writing, software 
 distributed under the License is distributed on an "AS IS" BASIS, 
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. 
 See the License for the specific language governing permissions and 
 limitations under the License.
-->
<?php
	$memcache = new Memcache;

	$info['name'] = "hometown";
	$info['city'] = "Philadelphia";
	$info['state'] = "PA";
	$info['latitude'] = "39.995664772727"; 
	$info['longitude'] = "-75.138788636364";

	$secret = json_encode($info,JSON_NUMERIC_CHECK);

	$memcache->set("secret2", $secret);
	$value = $memcache->get("secret2");	

	echo $value

		







?>