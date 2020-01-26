(function ($, $S) {
      //$ jQuery
      //$S window.localStorage
      //Variables Declaration
      var $board = $('#board'),
          //Board where the StickyNotes are sticked
          StickyNote, //Singleton Object containing the Functions to work with the LocalStorage
          len = 0, //Length of Objects in the LocalStorage 
          currentNotes = '', //Storage the html construction of the notes
          o; //Actual StickyNote data in the localStorage
   
   
   
      //Manage the StickyNotes in the Local Storage
        //Each note is saved in the localStorage as an Object  
      StickyNote = {
          add: function (obj) {
              obj.id = $S.length;
              $S.setItem(obj.id, JSON.stringify(obj));
          },
   
          retrive: function (id) {
              return JSON.parse($S.getItem(id));
          },
   
          remove: function (id) {
              $S.removeItem(id);
          },
   
          removeAll: function () {
              $S.clear();
          }
   
      };
   
      //If exist any note, Create it/them
      len = $S.length;
      if (len) {
          for (var i = 0; i < len; i++) {
              //Create all notes saved in localStorage
              var key = $S.key(i);
              o = StickyNote.retrive(key);
              currentNotes += '<div class="note"';
              currentNotes += ' style="left:' + o.left;
              currentNotes += 'px; top:' + o.top;
                    //data-key is the attribute to know what item delete in the localStorage
              currentNotes += 'px"><div class="toolbar"><span class="delete" data-key="' + key;
              currentNotes += '">x</span></div><div contenteditable="true" class="editable">';
              currentNotes += o.text;
              currentNotes += '</div>';
          }
   
          //Append all the notes to the board
          $board.html(currentNotes);
      }
   
      //When the document is ready, make all notes Draggable
      $(document).ready(function () {
          $(".note").draggable({
              cancel: '.editable',
            "zIndex": 3000,
            "stack" : '.note'
          });
      });
   
      //Remove StickyNote
      $('span.delete').live('click', function () {
          if (confirm('Are you sure you want to delete this Note?')) {
              var $this = $(this);
                    //data-key is the attribute to know what item delete in the localStorage
              StickyNote.remove($this.attr('data-key'));
              $this.closest('.note').fadeOut('slow', function () {
                  $(this).remove();
              });
          }
      });
   
      //Create note
      $('#btn-addNote').click(function () {
          $board.append('<div class="note" style="left:20px;top:70px"><div class="toolbar"><span class="delete" title="Close">x</span></div><div contenteditable class="editable"></div></div>');
          $(".note").draggable({
              cancel: '.editable'
          });
      });
   
      //Save all the notes when the user leaves the page
      window.onbeforeunload = function () {
          //Clean the localStorage
          StickyNote.removeAll();
          //Then insert each note into the LocalStorage
              //Saving their position on the page, in order to position them when the page is loaded again
          $('.note').each(function () {
              var $this = $(this);
              StickyNote.add({
                  top: parseInt($this.position().top),
                  left: parseInt($this.position().left),
                  text: $this.children('.editable').text()
              });
          });
      }
  })(jQuery, window.localStorage);