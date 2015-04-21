/*
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
*/
myFirebase.on('child_added', function(snapshot) {
    var msg = snapshot.val();
    displayMessage(msg.name,msg.text);
    checkSelf();
});

function checkSelf(){
    var self = document.querySelector("#name").value;

    if (self.length == 0){
        return;
    }   
            
    [].forEach.call( document.querySelectorAll('.msg'), function(el) {
        if (el.innerHTML.indexOf('<div class="name">'+self) > -1){
               el.classList.add("self");        
         } else{
               el.classList.remove("self"); 
        }
    }, false);
}
