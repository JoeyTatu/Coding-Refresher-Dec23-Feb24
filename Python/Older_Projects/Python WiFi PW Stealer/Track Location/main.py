import phonenumbers
# import opencage
# import os
# import sys
import json
import folium

from phonenumbers import geocoder, carrier
from opencage.geocoder import OpenCageGeocode
from phone_number import number, api_key

# Get Location
pepnumbner = phonenumbers.parse(number)
location = geocoder.description_for_number(pepnumbner, "en")
if location == "":
    location = "Unknown"
print(location)

# Get proivder (Note: Will get original service provider, even when number ported.)
service_provider = phonenumbers.parse(number)
provider = carrier.name_for_number(service_provider, "en")
if provider == "":
    provider = "Unknown"
print(provider)

geocoder = OpenCageGeocode(api_key)
query = str(location)
results = geocoder.geocode(query)

outputfile = "geocoding_results.json"
with open(outputfile, "w") as json_file:
    json.dump(results, json_file, indent=2)

lat = results[0]['geometry']['lat']
lng = results[0]['geometry']['lng']
print(lat, lng)

myMap = folium.Map(location=[lat, lng], zoom_start=9)
folium.Marker([lat, lng], popup=location).add_to(myMap)
myMap.save("Phone_number_location.html")
