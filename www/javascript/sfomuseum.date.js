var sfomuseum = sfomuseum || {};

sfomuseum.date = (function(){

    var self = {

	'toEDTFString': function(date, on_success, on_error){

	    var params = {
		"date": date,
	    };
	    
	    self.call('/api/sfomuseum/to-edtf-string', params, on_success, on_error);
	},

	'toEDTFDate': function(date, on_success, on_error){

	    var params = {
		"date": date,
	    };
	    
	    self.call('/api/sfomuseum/to-edtf-date', params, on_success, on_error);
	},	

	'call': function(method, data, on_success, on_error){
    
	    var dothis_onsuccess = function(rsp){

		if (on_success){
		    on_success(rsp);
		}
	    };
	    
	    var dothis_onerror = function(rsp){
			
		if (on_error){
		    on_error(rsp);
		}
	    };

	    var endpoint = location.protocol + '//' + location.host;
	    var url = endpoint + method;

	    var params = new URLSearchParams()

	    for (key in data){
		params.append(key, data[key]);
	    }
	    
	    url = url + "?" + params.toString();
	    
	    console.log("CALL", url);

	    /*
	    var form_data = data;

	    if (! form_data.append){

		form_data = new FormData();
		
		for (key in data){
		    form_data.append(key, data[key]);
		}
	    }
	    */
	    
	    var onload = function(rsp){

		var target = rsp.target;

		if (target.readyState != 4){
		    return;
		}

		var status_code = target['status'];
		var status_text = target['statusText'];
		
		var raw = target['responseText'];
		var data = undefined;

		try {
		    data = JSON.parse(raw);
		}

		catch (e){
		    dothis_onerror("Failed to parse JSON " + e);
		    return false;
		}

		dothis_onsuccess(data);
		return true;
	    };
	    
	    var onprogress = function(rsp){
		// console.log("progress");
	    };

	    var onfailed = function(rsp){

		dothis_onerror("Connection failed " + rsp);
	    };

	    var onabort = function(rsp){

		dothis_onerror("Connection aborted " + rsp);
	    };

	    // https://developer.mozilla.org/en-US/docs/Web/API/XMLHttpRequest/Sending_and_Receiving_Binary_Data

	    try {
		var req = new XMLHttpRequest();

		req.addEventListener("load", onload);
		req.addEventListener("progress", onprogress);
		req.addEventListener("error", onfailed);
		req.addEventListener("abort", onabort);

		/*
		for (var pair of form_data.entries()){
			console.log(pair[0]+ ', '+ pair[1]); 
		}
		*/
		
		req.open("GET", url, true);
		req.send();
		
		// req.send(form_data);
		
	    } catch (e) {
		
		dothis_onerror("Failed to send request, because " + e);
		return false;
	    }

	    return false;
	},

	
    };

    return self;
})();
