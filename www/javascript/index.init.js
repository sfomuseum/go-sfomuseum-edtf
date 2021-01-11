window.addEventListener("load", function load(event){

    var submit_button = document.getElementById("submit");

    if (! submit_button){
	console.log("Missing submit button");
	return;
    }

    submit_button.onclick = function(){

	var date_el = document.getElementById("date");

	if (! date_el){
	    console.log("Missing date el");
	    return false;
	}

	var date_str = date_el.value;

	if (date_str == ""){
	    console.log("Empty date string");
	    return false;
	}

	console.log("PARSE", date_str);
	
	to_edtf_string(date_str);
	to_edtf_date(date_str);
	
	return false;
    };

    submit_button.removeAttribute("disabled");
});
