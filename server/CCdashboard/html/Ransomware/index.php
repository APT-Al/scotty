<?php
include_once "../base.html";
include_once "leftpanel.html";


function getAllVictims(){
    # validate input for sql injection
    # end points should response only local requests except few of them
    $infos = file_get_contents("http://138.68.86.190:8080/dashboard/status/getallvictims");
    //$infos = file_get_contents("http://127.0.0.1/ok.json");
    return json_decode($infos);
}

function createVictimTable(){
    $victims = getAllVictims();
    $table = '<table class="table table-bordered">';
    $table.= '<thead>
                    <tr>
                    <th scope="col">ID</th>
                    <th scope="col">IP</th>
                    <th scope="col">Username</th>
                    <th scope="col">Infection Date</th>
                    <th scope="col">First Touch Date</th>
                    <th scope="col">Country</th>
                    </tr>
                </thead>
                <tbody>';
    
    foreach($victims as $vic){
        $table.='<tr>';
        $table.='<td>'.htmlspecialchars($vic->id).'</td>';
        $table.='<td>'.htmlspecialchars($vic->ip).'</td>';
        $table.='<td>'.htmlspecialchars($vic->username).'</td>';
        $table.='<td>'.htmlspecialchars($vic->infection_date).'</td>';
        $table.='<td>'.htmlspecialchars($vic->first_touch).'</td>';
        $table.='<td>'.htmlspecialchars($vic->country).'</td>';
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
        var url = "http://138.68.86.190:8080/dashboard/status/worldmap-ransomware";
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
        <?php print(createVictimTable());?>
    </div>




    
    


</main>