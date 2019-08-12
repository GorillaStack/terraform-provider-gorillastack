## 0.2.5 (August 12, 2019)

#### BREAKING CHANGES:
#### FEATURES:
#### IMPROVEMENTS:
#### BUG FIXES:

* Adding custom JSON unmarshal function to help construct StringArrayOrNull from API responses

#### NOTES:

* Version bumped to version 0.2.5 for this bugfix


## 0.2.4 (August 9, 2019)

#### BREAKING CHANGES:
#### FEATURES:
#### IMPROVEMENTS:
#### BUG FIXES:

* Fixing issue where empty array of email addresses was sent to the GorillaStack API where no extra email recipients desired. Property needed to be omitted from the request.

#### NOTES:

* Version bumped to version 0.2.4 for this bugfix


## 0.2.3 (August 5, 2019)

#### BREAKING CHANGES:
#### FEATURES:
#### IMPROVEMENTS:

* Adding tests around the StringArrayOrNull stringification
* Adding tests around the construction of the context from resource data

#### BUG FIXES:

* Fixing issue whereby `account_ids` would be interpreted as null, even if there were `account_group_ids` specified.

#### NOTES:

* Version bumped to version 0.2.3 for this bugfix


## 0.2.2 (August 2, 2019)

#### BREAKING CHANGES:
#### FEATURES:
#### IMPROVEMENTS:
#### BUG FIXES:

* Rule Resource was not passing through the user defined user_group attribute

#### NOTES:

* Version bumped to version 0.2.2 for this bugfix


## 0.2.1 (July 31, 2019)

#### BREAKING CHANGES:
#### FEATURES:
#### IMPROVEMENTS:
#### BUG FIXES:

* Fixing an infinite loop when stringifying the StringArrayOrNull type in the context

#### NOTES:

* Version bumped to version 0.2.1 for this bugfix


## 0.2.0 (July 26, 2019)

#### BREAKING CHANGES:

* CloudTrail triggers must specify match_fields as a single list item rather than a map. There was an oversight regarding the values of the map being arrays of strings, while Terraform maps only support strings

#### FEATURES:


* **New Trigger Type for Rule Resources:** CloudTrail Event
* **New Action Type for Rule Resources:** CheckTagCompliance
* **New Action Type for Rule Resources:** NotifyEvent (for CloudWatch Events)

#### IMPROVEMENTS:

* Improving toggling of adding the API version to the path so this doesn't interfere with local development when the API destination is overriden.

#### BUG FIXES:

* Fixing broken implementation of the CloudTrail trigger action (see notes in breaking changes above)

#### NOTES:

* Version bumped to version 0.2.0 as there are some minor breaking changes and significant new features, thus a patch version is not quite appropriate.


## 0.1.0 (July 8, 2019)

#### BREAKING CHANGES:

* None

#### FEATURES:

* **New Resource:** Rule
* **New Resource:** Tag Group

#### IMPROVEMENTS:
#### BUG FIXES:
#### NOTES:

* This is the initial GorillaStack Terraform provider implementation. As implied by a < 1 major version, there is little commitment to maintaining backwards compatibility right now, and there may be breaking changes in coming minor version updates. All breaking changes will be communicated in the changelog.
