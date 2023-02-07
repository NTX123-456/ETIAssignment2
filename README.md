# Documentation

<h1> Concept of Application: Travel Planning/Booking Application </h1>

<h2> Description </h2>
<p> A simple travel planning application which users can use to plan their travel and itinerary </p>

<h1> Features </h1>


<h2> Planning of itinerarie </h2>
<h3> Description: Within the console of the option 'Planning of itinerarie' is an details plan of journey that is associated with features like viewing, creating, and updating of your personal itinerarie. </h3>
<h3> Options </h3>
<h4> 1. View of itinerarie </h4>
<p> This feature is a detailed list of places that users have created into the itinerarie that consists of information like Location, Duration(days), Start Date, and End Date. With all this data results it could be retrieved through MYSQL database with the SELECT statement. </p>

<h4> 2. Create of itinerarie </h4>
<p> This feature will enables a outline of users creating travel itinerary through a new records into MYSQL database with the INSERT statement, that could includes adding of details of Location, Duration(days), Start Date, and End Date. 

Obtaining users question like:
1. Please enter the location:
2. Please enter the duration of travel (days):
3. Please enter start date (dd/mm/yyyy):
4. Please enter end date (dd/mm/yyyy):
</p>

<h4> 3. Update itinerarie </h4>
<p> This feature designed to allows users to make amend changes in duration of travel, start date, and end date. Subsequently, with the SQL UPDATE statement it will execute the modification towards the exisitng records into the database of data Duration(days), Start Date, and End Date. 

Obtaining users question to be updated:
1. Please enter your location to be updated:
2. Please enter the duration of travel to be updated:
3. Please enter the start date to be updated (dd/mm/yyyy):
4. Please enter the end date to be updated (dd/mm/yyyy):
</p>




<h2> Weather Update </h2>




<h2> View Hotel Information </h2>
<h3> Description: A Feature that allows users to select and display a list of hotels available in Singapore and other countries </h3>
<h3> Options </h3>
<h4> 1. Listing All Hotels </h4>
<p> This feature allows the user to display all the hotels available in the database with no filters. </p>

<h4> 2. Display Hotels from a certain country </h4>
<p> This feature allows the user to input a certain country to display a hotel from. After the user has input a country they want, the console will display all the hotels from that particular country. 
  
Current countries available in the database:
1. Singapore
2. Malaysia
3. Thailand
4. Indonesia
5. Australia
</p>

<h4> 3. Display Hotels with a certain star </h4>
<p> This feature allows the user to display all the hotels with a certain star rating. For example, the user can input the number 4 into the console and it will retrieve all the 4 star hotels from the database and display it. </p>

<h4> 4. Display Hotels with certain amenities </h4>
<p> This feature allows the user to display all the hotels with certain amenities. For example, if the user wants to view hotels with Pools, the user can input pools and the console will retrieve all the hotels with pools in their amenities. </p>

<h4> 5. Display Hotels from a certain price range </h4>
<p> This feature allows the user to display hotels from a certain price range they want. There are a total of 5 price ranges the user can choose from

  1. Hotels below $50
  2. Hotels between $50 and $100
  3. Hotels between $100 and $200
  4. Hotels between $200 and $300
  5. Hotels between $300 and $400
</p>




<h2> Booking Hotels & Attractions </h2>
<h3> Description: A Feature that allows user to view a list of hotel and attraction prices, after which user can start to book them. User can also choose to update or cancel their booking. </h3>

<h3> Options </h3>
<h4> User will be required to input their name to begin with this feature. </h4>
<h4> 1. Listing all tourist and attractions and hotels </h4>
<p> This option allow user to list out all tourist attractions and hotels that are available with the pricing. </p>

<h4> 2. Listing all tourist and attractions and hotels </h4>
<p> This option allows user to book an attraction or hotel by choosing the correct id, once id is confirmed, click on enter and the booking will be successful. </p>
<p> Example of Attractions and Hotels are in Singapore’s context. </p>

<h4> 3. Cancel or update booking </h4>
<p> This option allows user to either cancel a booking by typing out ‘cancel <id>’ or update a booking by typing out ‘update <id>’, where id is the number located on the left of each attraction or hotel. 
For update of booking the console will prompt user to pick another attraction or hotel by keying in the ideal id, after which the booking will be updated successfully.
</p>

<h4> Navigation options </h4>
<p> Type ‘main’ to return to main menu. </p>
<p> Type ‘quit’ to exit the program. </p>
