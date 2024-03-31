# use cases
- trying to make a build, and want to see for each item category, what is the list of item other people have used, and how many people have used them
    - would like to select to filter by a certain version
    - for each item, sort by:
        - number of builds that has it
        - total number of likes
        - average win rate?
        - highest win rate?
- when clicking on an item, want to see all builds that include that item
    - this can show if this item is used with certain other items
    - sort by:
        - number of likes
        - win rate
        - date

# enhancements
- [ ] internal cache - not really need to continuously be pinging the api. can store data and only refresh from api once in a while

# current issues
- [x] for each item, need to parse it's item category from the description