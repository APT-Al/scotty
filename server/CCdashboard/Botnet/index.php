<?php

include_once "../base.html";
include_once "leftpanel.html";


function getAllAgents(){
    # validate input for sql injection
    # end points should response only local requests except few of them
    //$infos = file_get_contents("http://138.68.86.190:8080/dashboard/status/getallagents");
    $infos = file_get_contents("http://127.0.0.1/ok.json");
    return json_decode($infos);
}

function createAgentTable(){
    $victims = getAllAgents();
    $table = '<table class="table table-bordered">';
    $table.= '<thead>
                    <tr>
                        <th scope="col">ID</th>
                        <th scope="col">IP</th>
                        <th scope="col">Current Status</th>
                        <th scope="col">Last Appeared</th>
                        <th scope="col">Country</th>
                    </tr>
                </thead>
                <tbody>';
    
    foreach($victims as $vic){
        $table.='<tr>';
        foreach($vic as $ky => $vl){
            #echo $ky."=>".$vl."<br>";
            $table.='<td>'.$vl.'</td>';
        }
        $table.='</tr>';
    }

    $table.='</tbody>
            </table>';
    
    return $table;
}


?>

<main class="col-md-9 ms-sm-auto col-lg-11 px-md-4">

<script src="/js/core.js"></script>
    <script src="/js/maps.js"></script>
    <script src="/js/animated.js"></script>
    <script src="/js/worldLow.js"></script>
    <script src="/js/worldmap.js"></script>
    <script>
        var url = "http://138.68.86.190:8080/dashboard/status/worldmap-botnet";
        //var url = "/ok.json";
        var worldmapvalues=[];
        var xhr = new XMLHttpRequest();
        xhr.open("GET", url, true);
        //xhr.setRequestHeader("Content-type", "application/json");
        xhr.onreadystatechange = function() {
            if (this.readyState == 4 && this.status == 200) {
            // Typical action to be performed when the document is ready:
                var temp = JSON.parse(xhr.responseText);
                for( var k in temp){
                    worldmapvalues.push({id:k,value:temp[k]});
                }
                polygonSeries.data = worldmapvalues;
                polygonSeries.exclude = ["AQ"];
            }
        };
        xhr.send();
    </script>

    <div>-</div>

    <div id="chartdiv" style="height: 600px;"></div>

    <div class="table-responsive">
        <?php print(createAgentTable());?>
    </div>



</main>



<!--

{
    "0": {
        "id": "1",
        "ip": "192.168.1.1",
        "current_status": "okan",
        "infection_date": "19/12/2020 21:33",
        "country": "Turkey"
    },
    "1": {
        "id": "3",
        "ip": "12.32.12.1",
        "current_status": "ogulcan",
        "last_appeared": "11/11/2020 21:33",
        "country": "Turkey"
    },
    "2": {
        "id": "4",
        "ip": "34.43.15.12",
        "current_status": "idil",
        "last_appeared": "19/12/2020 11:33",
        "country": "Germany"
    }
}



    -->